package goshopify

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"gopkg.in/jarcoal/httpmock.v1"
)

func variantTests(t *testing.T, variant Variant) {
	// Check that the ID is assigned to the returned variant
	expectedInt := int64(1)
	if variant.ID != expectedInt {
		t.Errorf("Variant.ID returned %+v, expected %+v", variant.ID, expectedInt)
	}

	// Check that the Title is assigned to the returned variant
	expectedTitle := "Yellow"
	if variant.Title != expectedTitle {
		t.Errorf("Variant.Title returned %+v, expected %+v", variant.Title, expectedTitle)
	}

	expectedInventoryItemId := int64(1)
	if variant.InventoryItemId != expectedInventoryItemId {
		t.Errorf("Variant.InventoryItemId returned %+v, expected %+v", variant.InventoryItemId, expectedInventoryItemId)
	}
}

func TestVariantList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/products/1/variants.json", globalApiPathPrefix),
		httpmock.NewStringResponder(200, `{"variants": [{"id":1},{"id":2}]}`))

	variants, err := client.Variant.List(1, nil)
	if err != nil {
		t.Errorf("Variant.List returned error: %v", err)
	}

	expected := []Variant{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(variants, expected) {
		t.Errorf("Variant.List returned %+v, expected %+v", variants, expected)
	}
}

func TestVariantCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/products/1/variants/count.json", globalApiPathPrefix),
		httpmock.NewStringResponder(200, `{"count": 3}`))

	params := map[string]string{"created_at_min": "2016-01-01T00:00:00Z"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/products/1/variants/count.json", globalApiPathPrefix),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Variant.Count(1, nil)
	if err != nil {
		t.Errorf("Variant.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Variant.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Variant.Count(1, CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Variant.Count returned %d, expected %d", cnt, expected)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Variant.Count returned %d, expected %d", cnt, expected)
	}
}

func TestVariantGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/variants/1.json", globalApiPathPrefix),
		httpmock.NewStringResponder(200, `{"variant": {"id":1}}`))

	variant, err := client.Variant.Get(1, nil)
	if err != nil {
		t.Errorf("Variant.Get returned error: %v", err)
	}

	expected := &Variant{ID: 1}
	if !reflect.DeepEqual(variant, expected) {
		t.Errorf("Variant.Get returned %+v, expected %+v", variant, expected)
	}
}

func TestVariantCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/products/1/variants.json", globalApiPathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("variant.json")))

	price := decimal.NewFromFloat(1)

	variant := Variant{
		Option1: "Yellow",
		Price:   &price,
	}
	result, err := client.Variant.Create(1, variant)
	if err != nil {
		t.Errorf("Variant.Create returned error: %v", err)
	}
	variantTests(t, *result)
}

func TestVariantUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/variants/1.json", globalApiPathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("variant.json")))

	variant := Variant{
		ID:      1,
		Option1: "Green",
	}

	variant.Option1 = "Yellow"

	returnedVariant, err := client.Variant.Update(variant)
	if err != nil {
		t.Errorf("Variant.Update returned error: %v", err)
	}
	variantTests(t, *returnedVariant)
}

func TestVariantDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/products/1/variants/1.json", globalApiPathPrefix),
		httpmock.NewStringResponder(200, "{}"))

	err := client.Variant.Delete(1, 1)
	if err != nil {
		t.Errorf("Variant.Delete returned error: %v", err)
	}
}
