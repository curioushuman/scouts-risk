package fixtures

import "scouts-risk/internal/model"

func ActTypeActivities(actType model.ActivityType) []model.Activity {
	if (actType == model.ActivityTypeFormal) {
	return ActivitiesFormal
	}
	return ActivitiesInformal
}

var Activities = map[string]model.Activity{
"cookingFire": {
	Label: "Cooking on open fire",
  Id: "cookingFire",
	ActType: model.ActivityTypeFormal,
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

var ActivitiesInformal = []model.Activity{
{
	Label: "Swimming",
  Id: "swimming",
	ActType: model.ActivityTypeInformal,
	Hazards: []model.Hazard{
		{
			Label: "Insufficient ability",
			Id: "pool1",
			Consequences: []model.Consequence{
				{
					Label: "Drowning",
					Id: "pool1C1",
				},
			},
			Controls: []model.Control{
				{
					Label: "Sufficient safety briefing",
					Id: "pool12",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Buddy system",
					Id: "pool13",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Monitor scouts while in use",
					Id: "pool22",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "pool14",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "pool15",
					When: model.ControlWhenResponse,
				},
			},
		},
		{
			Label: "Improper use",
			Id: "pool2",
			Consequences: []model.Consequence{
				{
					Label: "Drowning",
					Id: "pool1C1",
				},
			},
			Controls: []model.Control{
				{
					Label: "Sufficient safety briefing",
					Id: "pool21",
					When: model.ControlWhenDuring,
				},
				{
					Label: "Monitor scouts while in use",
					Id: "pool22",
					When: model.ControlWhenDuring,
				},
				{
					Label: "First aid",
					Id: "pool23",
					When: model.ControlWhenResponse,
				},
				{
					Label: "Call ambulance",
					Id: "pool24",
					When: model.ControlWhenResponse,
				},
			},
		},
	},
},
}

var ActivitiesFormal = []model.Activity{
	Activities["cookingFire"],
}

var ActivitiesDuplicate = []model.Activity{
	Activities["cookingFire"],
	ActivitiesInformal[0],
}