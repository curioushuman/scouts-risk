// @ts-nocheck - so we can use HTMX attributes
function globalData() {
  const uap = new UAParser(navigator.userAgent);
  const browser = uap.getBrowser();
  const os = uap.getOS();
  return {
    // * Basic user agent
    uniqueId: '',
    browser: '',
    os: '',
    // * Feedback form
    feedbackOpen: false,
    feedbackDisplayHelp: true,
    openFeedback: function (resetFields) {
      resetFields.forEach((field) => {
        field.value = '';
      });
      this.feedbackOpen = true;
      this.uniqueId = [browser.major, os.version].join('-');
      this.browser = [browser.name, browser.major].join('-');
      this.os = [os.name, os.version].join('-');
    },
    closeFeedback: function () {
      this.feedbackOpen = false;
      this.feedbackDisplayHelp = false;
    },
  };
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

// A combo of severity and likelihood
const ratings = {
  'rare-insignificant': 'low',
  'rare-minor': 'low',
  'rare-moderate': 'low',
  'rare-significant': 'medium',
  'rare-severe': 'medium',
  // --
  'unlikely-insignificant': 'low',
  'unlikely-minor': 'low',
  'unlikely-moderate': 'medium',
  'unlikely-significant': 'medium',
  'unlikely-severe': 'high',
  // --
  'possible-insignificant': 'low',
  'possible-minor': 'medium',
  'possible-moderate': 'medium',
  'possible-significant': 'high',
  'possible-severe': 'high',
  // --
  'likely-insignificant': 'low',
  'likely-minor': 'medium',
  'likely-moderate': 'high',
  'likely-significant': 'high',
  'likely-severe': 'extreme',
  // --
  'almost_certain-insignificant': 'low',
  'almost_certain-minor': 'high',
  'almost_certain-moderate': 'high',
  'almost_certain-significant': 'extreme',
  'almost_certain-severe': 'extreme',
};

const badgeColour = {
  low: 'badge-success',
  medium: 'badge-info',
  high: 'badge-warning',
  extreme: 'badge-error',
};

function initHazards(initialConsequences) {
  return {
    ratings,
    consequences: initialConsequences,
    addConsequence: function (consequence) {
      this.consequences.push(consequence);
    },
    badgeColour: function (ratingCombo) {
      const classes = {};
      if (!ratings[ratingCombo]) {
        classes['badge-ghost'] = true;
      } else {
        classes[badgeColour[ratings[ratingCombo]]] = true;
      }
      return classes;
    },
  };
}
