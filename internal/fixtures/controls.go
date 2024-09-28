package fixtures

import "scouts-risk/internal/model"

var ControlsExisting = []model.Control{
	{
		Label: "All participants told to bring high boots, long pants",
		Id: "91",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Tree / branch inspection",
		Id: "22",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Ensure installation by skilled person",
		Id: "31",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Review / check equipment prior to use",
		Id: "32",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Make sure safety rails/precautions exist",
		Id: "41",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Sufficient safety briefing",
		Id: "12",
		When: model.ControlWhenDuring,
	},
	{
		Label: "Utilise necessary safety equipment; for height",
		Id: "42",
		When: model.ControlWhenDuring,
	},
	{
		Label: "Monitor scouts while in use",
		Id: "72",
		When: model.ControlWhenDuring,
	},
	{
		Label: "All participants to wear high boots, long pants in bushy area",
		Id: "92",
		When: model.ControlWhenDuring,
	},
	{
		Label: "Stay with/behind a leader where possible",
		Id: "93",
		When: model.ControlWhenDuring,
	},
	{
		Label: "Check the area prior to use",
		Id: "94",
		When: model.ControlWhenDuring,
	},
	{
		Label: "First aid",
		Id: "13",
		When: model.ControlWhenResponse,
	},
	{
		Label: "Call ambulance",
		Id: "35",
		When: model.ControlWhenResponse,
	},
	{
		Label: "Review / maintain equipment regularly",
		Id: "33",
		When: model.ControlWhenRegular,
	},
}

var ControlsBefore = map[string][]model.Control{
"treeHouse": {
	{
		Label: "Cover the sharp edges",
		Id: "1",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Tree / branch inspection",
		Id: "2",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Review / check equipment prior to use",
		Id: "2",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Make sure safety rails/precautions exist",
		Id: "1",
		When: model.ControlWhenBefore,
	},
},
"bushy": {
	{
		Label: "Tree / branch inspection",
		Id: "2",
		When: model.ControlWhenBefore,
	},
	{
		Label: "Bushfire inspection",
		Id: "2",
		When: model.ControlWhenBefore,
	},
},
"pack": {
	{
		Label: "All participants told to bring high boots, long pants",
		Id: "1",
		When: model.ControlWhenBefore,
	},
},
}