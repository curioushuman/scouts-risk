package fixtures

import "scouts-risk/internal/model"

var Activities = map[string]model.Activity{
"cookingFire": {
	Label: "Cooking on open fire",
  Id: "cookingFire",
	ActType: model.ActivityTypeInformal,
	Hazards: []model.Hazard{
		{
			Label: "Open fire",
			Id: "A1",
			Consequences: []model.Consequence{
				{
					Label: "Burn",
					Id: "A1C1",
				},
			},
			Controls: []model.Control{
				{
					Label: "Pack fire blanket and extinguisher",
					Id: "A11",
					When: model.ControlWhenBefore,
				},
				{
					Label: "Safety briefing",
					Id: "A12",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "A13",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "A14",
					When: model.ControlWhenResponse,
				},
			},
		},
	},
},
}