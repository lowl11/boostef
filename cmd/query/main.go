package main

import (
	"fmt"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/pkg/builder"
	"github.com/lowl11/boostef/pkg/enums/signs"
	"github.com/lowl11/boostef/pkg/query"
)

func main() {
	//insertQuery()
	//updateQuery()
	//deleteQuery()
	//selectQuery()
}

func insertQuery() {
	fmt.Println(builder.
		Insert([]query.Pair{
			{
				Column: "id",
				Value:  1,
			},
			{
				Column: "title_ru",
				Value:  "my title ru",
			},
			{
				Column: "title_kz",
				Value:  "my title kz",
			},
			{
				Column: "count",
				Value:  20,
			},
		}...).
		To("products").
		OnConflict("DO NOTHING").
		Get())
}

func updateQuery() {
	fmt.Println(builder.
		Update("products").
		Set([]query.Pair{
			{
				Column: "count",
				Value:  0,
			},
			{
				Column: "deleted",
				Value:  true,
			},
		}...).
		Where(func(where iquery.Where) {
			where.Equal("id", 5)
		}).
		Get())
}

func deleteQuery() {
	fmt.Println(builder.
		Delete("products").
		Where(func(where iquery.Where) {
			where.Gte("id", 5)
		}).
		Get())
}

func selectQuery() {
	fmt.Println(builder.
		Select("id", "username", "password", "product.title").
		From("users").
		SetAlias("user").
		Where(func(where iquery.Where) {
			where.
				Equal("age", 25).
				ILike("username", "%lazy_owl%")

			where.Not(func(where iquery.Where) iquery.Where {
				return where.
					Bool("is_resident", true).
					Or(func(where iquery.Where) iquery.Where {
						return where.
							Equal("last_name", "Ussayev").
							Equal("first_name", "Erik")
					})
			})

			where.
				In("id", []any{1, 2, 3}).
				Between("title", "John", "Erik")
		}).
		OrderBy("title", "last_name", "first_name").
		Having(func(aggregate iquery.Aggregate) {
			aggregate.Count("id", signs.Gte, 250)
		}).
		//GroupBy("product.brand_id", "product.title_ru").
		GroupByAggregate(func(aggregate iquery.Aggregate) {
			aggregate.Avg("some_field", signs.Equal, true)
		}).
		Page(10, 1).
		Get())
}
