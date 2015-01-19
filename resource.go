package framgo

import (
	"encoding/json"
	"errors"
	"net/url"
)

type Resource struct {
	// Used to set response URLs as, CSS, JS or common links
	Links map[string][]string
	// Data to encode or pass as template input
	Data map[string]interface{}
	// Plain text
	Plain []byte
	// Content Type
	Content string
}

// Creates new resource
func NewResource() *Resource {
	var res Resource
	// 0 bytes
	res.Plain = make([]byte, 0)
	// map of links
	res.Links = make(map[string][]string)
	res.Data = make(map[string]interface{})
	res.Content = "text/plain"
	return &res
}

// Merge 2 resources
func (res *Resource) Merge(dataKey string, r *Resource) {
	defer func() {
		if e := recover(); e != nil {
			return
		}
	}()
	// mix all links
	if r.Links != nil && len(r.Links) > 0 {
		for k, v := range r.Links {
			for _, s := range v {
				res.AddLink(k, s)
			}
		}
	}
	// set data in new key
	if res.Data != nil {
		_, ok := res.Data[dataKey]
		if !ok {
			res.Data[dataKey] = r.Data
		}
	}
}

// Add links
func (res *Resource) AddLink(key string, val string) error {
	_, err := url.Parse(val)
	if err != nil {
		return err
	}

	_, ok := res.Links[key]
	if !ok {
		res.Links[key] = make([]string, 0)
	}

	res.Links[key] = append(res.Links[key], val)
	return nil
}

// Set data in key
func (res *Resource) AddData(k string, i interface{}) error {
	_, ok := res.Data[k]
	if !ok {
		res.Data[k] = i
		return nil
	} else {
		return errors.New("Key: " + k + " already used!")
	}
}

func (res *Resource) JSON(i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	res.Plain = b
	res.Content = "application/json"
	return nil
}
