package dup

import "github.com/tylland/transition/cdas/internal/infrastructure/dup"

type PlanningDataConsumer struct {
	query *dup.Query
}

func NewPlanningDataConsumer() *PlanningDataConsumer {
	filter := dup.EQ("", "")

	query := dup.NewQuery("OperativFardplan").
		SchemaVersion("3.0").
		Limit(20).Filter(filter)

	dup.NewRequest("").Query(query)

	return &PlanningDataConsumer{query: query}
}
