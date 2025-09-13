package repositories

import "errors"

var ErrNotFound = errors.New("Kayıt bulunamadı")

var ErrForbidden = errors.New("Bu işlemi yapmaya yekiniz yok")
