// @ts-nocheck
function globalData() {
  return {};
}

function initRegister() {
  return {
    // * Modals
    // Locations
    areaNewOpen: false,
    equipmentNewOpen: false,
    // Activities
    formalNewOpen: false,
    informalNewOpen: false,
    // Hazards
    hazardNewOpen: false,
    consequenceNewOpen: false,
    controlNewOpen: false,
    hazardModalOpen: function (targetId, form) {
      form.setAttribute('hx-target', targetId);
      form.reset();
      this.hazardNewOpen = true;
    },
    consequenceModalOpen: function (targetId, form, field) {
      form.setAttribute('hx-target', targetId);
      form.reset();
      field.setAttribute('value', targetId);
      this.consequenceNewOpen = true;
    },
    controlModalOpen: function (targetId, form, field) {
      form.setAttribute('hx-target', targetId);
      form.reset();
      field.setAttribute('value', targetId);
      this.controlNewOpen = true;
    },
  };
}

function initHazards(initialConsequences) {
  return {
    consequences: initialConsequences,
    addConsequence: function (consequence) {
      this.consequences.push(consequence);
    },
  };
}
