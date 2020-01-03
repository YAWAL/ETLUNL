package transform

import (
	"strconv"

	"github.com/YAWAL/ETLUNL/model"
)

// Transform
func Transform(rawData []model.RawData) (preparedData []model.PreparedData, err error) {

	for _, rawItem := range rawData {
		preparedItem, err := prepareItem(rawItem)
		if err != nil {
			// log err
			return preparedData, nil
		}
		preparedData = append(preparedData, preparedItem)
	}

	return preparedData, nil
}

func prepareItem(rawItem model.RawData) (preparedItem model.PreparedData, err error) {
	preparedItem.Date = rawItem.Date
	preparedItem.Lototron = rawItem.Lototron
	preparedItem.Game, err = strconv.Atoi(rawItem.Game)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.BallSet, err = strconv.Atoi(rawItem.BallSet)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball1, err = strconv.Atoi(rawItem.Ball1)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball2, err = strconv.Atoi(rawItem.Ball2)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball3, err = strconv.Atoi(rawItem.Ball3)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball4, err = strconv.Atoi(rawItem.Ball4)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball5, err = strconv.Atoi(rawItem.Ball5)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball6, err = strconv.Atoi(rawItem.Ball6)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball2Winners, err = strconv.Atoi(rawItem.Ball2Winners)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball3Winners, err = strconv.Atoi(rawItem.Ball3Winners)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball4Winners, err = strconv.Atoi(rawItem.Ball4Winners)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball5Winners, err = strconv.Atoi(rawItem.Ball5Winners)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball6Winners, err = strconv.Atoi(rawItem.Ball6Winners)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball2Price, err = strconv.Atoi(rawItem.Ball2Price)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball3Price, err = strconv.Atoi(rawItem.Ball3Price)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball4Price, err = strconv.Atoi(rawItem.Ball4Price)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball5Price, err = strconv.Atoi(rawItem.Ball5Price)
	if err != nil {
		return preparedItem, err
	}
	preparedItem.Ball6Price, err = strconv.Atoi(rawItem.Ball6Price)
	if err != nil {
		return preparedItem, err
	}

	return preparedItem, nil
}
