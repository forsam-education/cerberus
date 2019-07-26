/**
 * ROUTER PART
 */
// Used for the routes title
export const SiteName = 'Cerberus';
export const TitleSeparator = ' - ';

// Vue Router Mode config
export const RouterMode = 'history';

/**
 * API PART
 */
export const APIConfig = {
  baseURL: '/api/',
  withCredentials: true,
  crossDomain: true,
  contentType: false,
  responseType: 'json',
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  },
};

/**
 * GOOGLE ANALYTICS PART
 */
export const enableAnalytics = false;
export const analyticsKey = 'UA-XXX-X';
export const disableAnalyticsInDebug = true;
