// Code generated by Validator v0.1.4. DO NOT EDIT.

package favorite

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *DouyinFavoriteActionRequest) IsValid() error {
	if p.VideoId <= int64(0) {
		return fmt.Errorf("field VideoId gt rule failed, current value: %v", p.VideoId)
	}
	_src := []int8{int8(1), int8(2)}
	var _exist bool
	for _, src := range _src {
		if p.ActionType == int8(src) {
			_exist = true
			break
		}
	}
	if !_exist {
		return fmt.Errorf("field ActionType in rule failed, current value: %v", p.ActionType)
	}
	return nil
}
func (p *DouyinFavoriteActionResponse) IsValid() error {
	return nil
}
func (p *DouyinFavoriteListRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	return nil
}
func (p *DouyinFavoriteListResponse) IsValid() error {
	return nil
}
func (p *Video) IsValid() error {
	if p.Author != nil {
		if err := p.Author.IsValid(); err != nil {
			return fmt.Errorf("filed Author not valid, %w", err)
		}
	}
	return nil
}
func (p *UserInfo) IsValid() error {
	return nil
}
