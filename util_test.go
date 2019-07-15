package goshopify

import (
	"fmt"
	"testing"
)

func TestShopFullName(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"myshop", "myshop.myshopify.com"},
		{"myshop.", "myshop.myshopify.com"},
		{" myshop", "myshop.myshopify.com"},
		{"myshop ", "myshop.myshopify.com"},
		{"myshop \n", "myshop.myshopify.com"},
		{"myshop.myshopify.com", "myshop.myshopify.com"},
	}

	for _, c := range cases {
		actual := ShopFullName(c.in)
		if actual != c.expected {
			t.Errorf("ShopFullName(%s): expected %s, actual %s", c.in, c.expected, actual)
		}
	}
}

func TestShopShortName(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"myshop", "myshop"},
		{"myshop.", "myshop"},
		{" myshop", "myshop"},
		{"myshop ", "myshop"},
		{"myshop \n", "myshop"},
		{"myshop.myshopify.com", "myshop"},
		{".myshop.myshopify.com.", "myshop"},
	}

	for _, c := range cases {
		actual := ShopShortName(c.in)
		if actual != c.expected {
			t.Errorf("ShopShortName(%s): expected %s, actual %s", c.in, c.expected, actual)
		}
	}
}

func TestShopBaseUrl(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"myshop", "https://myshop.myshopify.com"},
		{"myshop.", "https://myshop.myshopify.com"},
		{" myshop", "https://myshop.myshopify.com"},
		{"myshop ", "https://myshop.myshopify.com"},
		{"myshop \n", "https://myshop.myshopify.com"},
		{"myshop.myshopify.com", "https://myshop.myshopify.com"},
	}

	for _, c := range cases {
		actual := ShopBaseUrl(c.in)
		if actual != c.expected {
			t.Errorf("ShopBaseUrl(%s): expected %s, actual %s", c.in, c.expected, actual)
		}
	}
}

func TestMetafieldPathPrefix(t *testing.T) {
	cases := []struct {
		resource   string
		resourceID int64
		expected   string
	}{
		{"", 0, fmt.Sprintf("%s/metafields", globalApiPathPrefix)},
		{"products", 123, fmt.Sprintf("%s/products/123/metafields", globalApiPathPrefix)},
	}

	for _, c := range cases {
		actual := MetafieldPathPrefix(c.resource, c.resourceID)
		if actual != c.expected {
			t.Errorf("MetafieldPathPrefix(%s, %d): expected %s, actual %s", c.resource, c.resourceID, c.expected, actual)
		}
	}
}

func TestFulfillmentPathPrefix(t *testing.T) {
	cases := []struct {
		resource   string
		resourceID int64
		expected   string
	}{
		{"", 0, fmt.Sprintf("%s/fulfillments", globalApiPathPrefix)},
		{"orders", 123, fmt.Sprintf("%s/orders/123/fulfillments", globalApiPathPrefix)},
	}

	for _, c := range cases {
		actual := FulfillmentPathPrefix(c.resource, c.resourceID)
		if actual != c.expected {
			t.Errorf("FulfillmentPathPrefix(%s, %d): expected %s, actual %s", c.resource, c.resourceID, c.expected, actual)
		}
	}
}
