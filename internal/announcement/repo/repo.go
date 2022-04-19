package store

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/jkrus/Test_Seller/pkg/models"
)

type (
	db struct {
		orm *gorm.DB
	}
)

func NewAnnouncementStorage(orm *gorm.DB) Storage {
	return &db{orm: orm}
}

func (d *db) Create(bytes []byte) (int64, error) {
	announcement, err := decode(bytes)
	if err != nil {
		return 0, err
	}

	err = validate(announcement)
	if err != nil {
		return 0, err
	}

	err = d.orm.Create(&announcement).Error
	if err != nil {
		return 0, err
	}

	return announcement.ID, nil
}

func (d *db) Get(id int64, opts []string) ([]byte, error) {
	m := &models.Announcement{}

	err := d.orm.Model(m).Where("id=?", id).Find(m).Error
	if err != nil {
		return nil, err
	}
	if m.ID == 0 {
		return nil, errors.New("id not found")
	}

	rez := &models.AnnouncementDTO{
		Price: m.Price,
		Name:  m.Name,
	}
	imgs := &models.ImagesDTO{Images: map[string]struct{}{}}
	for _, opt := range opts {
		switch opt {
		case "description":
			rez.Description = m.Description
		case "images":
			err = json.Unmarshal(m.Images, &imgs.Images)
			if err != nil {
				return nil, err
			}
			for val := range imgs.Images {
				rez.Images = append(rez.Images, val)
			}
		}
	}

	return encode(rez)
}

func (d *db) GetList(page, limit int, sortBy, sort string) ([]byte, error) {
	announcements := make([]*models.Announcement, 0, limit)
	offset := (page - 1) * limit
	var err error
	if page != 0 {
		if sortBy != "" {
			if sortBy == "data" {
				sortBy = "created_at"
			}
			err = d.orm.Limit(limit).Offset(offset).Order(sortBy + " " + sort).Find(&announcements).Error
		} else {
			err = d.orm.Limit(limit).Offset(offset).Find(&announcements).Error
		}
	} else {
		err = d.orm.Find(&announcements).Error
	}

	if err != nil {
		return nil, err
	}

	rez := make([]*models.AnnouncementDTO, 0, len(announcements))

	for _, an := range announcements {
		rez = append(rez, &models.AnnouncementDTO{Name: an.Name, Price: an.Price, Images: []string{an.MainImage}})
	}

	return encode(rez)
}

func (d *db) Update(id int, path string) ([]byte, error) {
	m := &models.Announcement{}
	err := d.orm.Model(m).Where("id=?", id).Find(m).Error
	if err != nil {
		return nil, err
	}

	imgs := models.ImagesDTO{Images: map[string]struct{}{}}
	if m.Images != nil {
		err = json.Unmarshal(m.Images, &imgs.Images)
		if err != nil {
			return nil, err
		}
	}
	imgs.Images[path] = struct{}{}
	bytes, err := json.Marshal(imgs.Images)
	if err != nil {
		return nil, err
	}
	if m.MainImage == "" {
		err = d.orm.Model(m).
			Where("id = ?", m.ID).
			Update("images", bytes).
			Update("main_image", path).
			Error
	} else {
		err = d.orm.Model(m).
			Where("id = ?", m.ID).
			Update("images", bytes).
			Error
	}
	if err != nil {
		return nil, err
	}

	return encode(m)
}

func decode(data []byte) (*models.Announcement, error) {
	m := &models.Announcement{}

	err := json.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func encode(v interface{}) ([]byte, error) {
	switch v.(type) {
	case *models.AnnouncementDTO:
		bytes, err := json.Marshal(v.(*models.AnnouncementDTO))
		if err != nil {
			return nil, err
		}
		return bytes, nil

	case []*models.AnnouncementDTO:
		bytes, err := json.Marshal(v.([]*models.AnnouncementDTO))
		if err != nil {
			return nil, err
		}
		return bytes, nil

	}

	return nil, nil
}

func validate(a *models.Announcement) error {
	if a.Name == "" {
		return errors.New("name can not be empty")
	}
	if a.Description == "" {
		return errors.New("description can not be empty")
	}

	if a.Price == 0 || a.Price < 0 {
		return errors.New("wrong price")
	}

	return nil
}
