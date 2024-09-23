package fixtures

import "scouts-risk/internal/model"

var Activities = map[string]model.Activity{
"cookingFire": {
	Label: "Cooking on open fire",
  Id: "cookingFire",
	Hazards: []model.Hazard{
		{
			Label: "Open fire",
			Id: "1",
			Controls: []model.Control{
				{
					Label: "Pack fire blanket and extinguisher",
					Id: "11",
					When: model.ControlWhenBefore,
				},
				{
					Label: "Safety briefing",
					Id: "12",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "13",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "14",
					When: model.ControlWhenResponse,
				},
			},
		},
	},
},
}