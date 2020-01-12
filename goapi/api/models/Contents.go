package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Contents struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title string `gorm:"size:255;not null; unique" json:"title"`
	Content string `gorm:"size:255; not null; unique" json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Contents) Prepare(){
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Contents) Validate() error{
	if p.Title == ""{
		return errors.New("Required Title")
	}
	if p.Content == ""{
		return errors.New("Required Content")
	}
	return nil
}

func (p *Contents) SavePost(db *gorm.DB)(*Contents, error){
	return p, nil
}

func (p *Contents) FindAllPosts(db *gorm.DB)(*Contents, error){
	return p, nil
}

func (p *Contents) FindPostByID(db *gorm.DB, pid uint64)(*Contents, error)  {
	return p, nil
}

func (p *Contents) UpdatePost(db *gorm.DB) (*Contents, error){
	return p, nil
}

func (p *Contents) DeletePost(db *gorm.DB, pid uint64)(int64, error){
	return db.RowsAffected, nil
}

