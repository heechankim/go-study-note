package main

import (
	"net/http"
	"strings"
)

type router struct {
	// key: http method
	// value: HandlerFunc executed by URL pattern
	handlers map[string]map[string]http.HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc) {
	// http 메서드로 등록된 맵이 있는지 확인
	m, ok := r.handlers[method]
	if !ok {
		// 등록된 맵이 없으면 새 맵을 생성
		m = make(map[string]http.HandlerFunc)
		r.handlers[method] = m
	}
	// http 메서드로 등록된 맵에 URL 패턴과 핸들러 함수 등록
	m[pattern] = h
}

func match(pattern, path string) (bool, map[string]string) {
	if pattern == path {
		return true, nil
	}

	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) != len(paths) {
		return false, nil
	}

	params := make(map[string]string)

	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
			// 패턴과 패스의 부분 문자열이 일치하면 바로 다음 루프 수행
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			// 패턴이 ':' 문자로 시작하면 params에 URL params 를 담은 후 다음 루프 수행
			params[patterns[i][1:]] = paths[i]
		default:
			// 일치하는 경우가 없으면 false 를 반환
			return false, nil
		}
	}
	return true, params
}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for pattern, handler := range r.handlers[req.Method] {
		if ok, _ := match(pattern, req.URL.Path); ok {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
	return
}
