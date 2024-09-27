package fixtures

import "scouts-risk/internal/model"

var HazardsFirstAid = []model.Hazard{
	{
		Label: "Exposed sharp edges / points",
		Id: "1",
		Consequences: []model.Consequence{
			{
				Label: "Tetanus",
				Id: "1C1",
			},
		},
		Controls: []model.Control{
			{
				Label: "Cover the sharp edges",
				Id: "11",
				When: model.ControlWhenBefore,
			},
			{
				Label: "Sufficient safety briefing",
				Id: "12",
				When: model.ControlWhenDuring,
			},
			{
				Label: "First aid",
				Id: "13",
				When: model.ControlWhenResponse,
			},
		},
	},
	{
		Label: "Loss of structural integrity",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Falling from height",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Falling branches",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Equipment failure",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Improper use",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Incorrect configuration",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Snakes",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Spiders",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Leeches",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Other fauna",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Insects",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Ticks",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Outdoor trip hazards",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Falling trees",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Falling branches",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Falling branches",
		Id: "1",
		Controls: []model.Control{},
	},
	{
		Label: "Falling branches",
		Id: "1",
		Controls: []model.Control{},
	},
}