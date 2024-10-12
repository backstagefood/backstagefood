# backstagefood
Backstage Food is an advanced restaurant order management system that empowers customers to seamlessly place their own orders. It offers a user-friendly interface for selecting menu items, customizing orders, and completing transactions through self-checkout, streamlining the dining experience for both customers and restaurant staff.

### Environment variables
- before start, you should create a .env defining the following variables:   
  SERVER_PORT   
  DB_USER   
  DB_PASS   

### Building and running
- execute the following steps:
> make
> docker compose up -d

### Links 
- [miro dashboard](https://miro.com/app/board/uXjVKg5JFS0=/)
- [swagger doc](http://localhost:8080/swagger/index.html)


> swag init --parseInternal --dir cmd/app/,internal/handlers/ --output internal/app/docs/


```


query := "SELECT id, id_category, description, ingredients, price, created_at, updated_at FROM products"
	params := make(map[string]string)
	if len(category) > 0 {
		params[category] = "id_category = ?"
	}
	if len(name) > 0 {
		params[name] = "name like '%?%'"
	}
	log.Printf("[connection] ListAllProducts %s %s, params=%s", category, name)

	if len(params) > 0 {
		query = query + " WHERE "
		for k, v := range params {
			query = query
			log.Printf("[connection] ListAllProducts parametro=%s, query=%s", k, v)
		}
	}

			category := c.QueryParam("category")
    		name := c.QueryParam("name")
    		products, err := uc.GetProducts(category, name)
```