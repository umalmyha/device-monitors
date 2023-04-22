package main

import (
	"gorm.io/gen"

	"github.com/umalmyha/device-monitors/devices-service/internal/model"
)

type DeviceQuerier interface {
	// SELECT id, name, description, latitude, longitude, updated_at, created_at FROM @@table WHERE name=@name
	FindByName(name string) (*gen.T, error)

	// SELECT id, name, description, latitude, longitude, updated_at, created_at FROM @@table
	//  {{where}}
	//  	{{if query.Search != nil && *query.Search != ""}}
	//			INSTR(name, @query.Search) > 0 OR INSTR(description, @query.Search) > 0
	//		{{end}}
	//      {{if query.FromLatitude != nil && query.ToLatitude != nil}}
	//			AND (latitude >= @query.FromLatitude AND latitude <= @query.ToLatitude)
	//		{{else if query.FromLatitude != nil}}
	//			AND latitude >= @query.FromLatitude
	//		{{else}}
	//			AND latitude <= @query.ToLatitude
	//		{{end}}
	//      {{if query.FromLongitude != nil && query.ToLongitude != nil}}
	//			AND (longitude >= @query.FromLongitude AND longitude <= @query.ToLongitude)
	//		{{else if query.FromLatitude != nil}}
	//			AND longitude >= @query.FromLongitude
	//		{{else}}
	//			AND longitude <= @query.ToLongitude
	//		{{end}}
	//  {{end}}
	FindAll(query model.GetAllDevicesQuery) ([]*gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./internal/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	g.ApplyBasic(&model.Device{})
	g.ApplyInterface(func(DeviceQuerier) {}, &model.Device{})
	g.Execute()
}
