package modelutil

import (
	"apigateway/gen/go/commonpb"
	"shared/models/commonmodel"
	"sort"
)

func BuildStatuses() []*commonmodel.StatusItem {
	// Allocate exact size
	statusesList := make([]*commonmodel.StatusItem, 0, len(commonpb.Status_value))

	for name, val := range commonpb.Status_value {
		statusesList = append(statusesList, &commonmodel.StatusItem{
			Name:  name,
			Value: val,
		})
	}

	// optional but recommended: stable order
	sort.Slice(statusesList, func(i, j int) bool {
		return statusesList[i].Value < statusesList[j].Value
	})

	return statusesList
}

func BuildPageTypes() []*commonmodel.PageTypeItem {
	// Allocate exact size
	pageTypeList := make([]*commonmodel.PageTypeItem, 0, len(commonpb.Status_value))

	for name, val := range commonpb.PageType_value {
		pageTypeList = append(pageTypeList, &commonmodel.PageTypeItem{
			Name:  name,
			Value: val,
		})
	}

	// optional but recommended: stable order
	sort.Slice(pageTypeList, func(i, j int) bool {
		return pageTypeList[i].Value < pageTypeList[j].Value
	})

	return pageTypeList
}

func BuildAdsPlatforms() []*commonmodel.AdsPlatformItem {
	// Allocate exact size
	adsPlatformList := make([]*commonmodel.AdsPlatformItem, 0, len(commonpb.Status_value))

	for name, val := range commonpb.PageType_value {
		adsPlatformList = append(adsPlatformList, &commonmodel.AdsPlatformItem{
			Name:  name,
			Value: val,
		})
	}

	// optional but recommended: stable order
	sort.Slice(adsPlatformList, func(i, j int) bool {
		return adsPlatformList[i].Value < adsPlatformList[j].Value
	})

	return adsPlatformList
}
