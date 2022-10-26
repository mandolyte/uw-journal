# dcs-js

Primary repos:
- https://github.com/unfoldingWord/dcs-js
- https://github.com/unfoldingWord/dcs-react-hooks

Also a set of hooks: :
- https://github.com/unfoldingWord/dcs-react-hooks

Repository API:
https://github.com/unfoldingWord/dcs-js/blob/master/apis/repository-api.ts


Markdown docs:
https://github.com/unfoldingWord/dcs-js/blob/master/documentation/classes/RepositoryApi.md#-repogetbranch

getApiConfig.js
```js
export const getApiConfig = ({ token, basePath = "https://qa.door43.org/api/v1/", ...config }) => ({
  apiKey: token && ((key) => key === "Authorization" ? `token ${token}` : ""),
  basePath: basePath?.replace(/\/+$/, ""),
  ...config
})
```
useRepoClient.js
```js
import PropTypes from 'prop-types';
import { RepositoryApi } from 'dcs-js';
import { AxiosInstance } from "axios";
import { getApiConfig } from "@helpers/api";

/**
 * Uses RepositoryApi from dcs-js.
 */
export const useRepoClient = ({ token, basePath, repoClient, axios, configuration } = {}) => {
  if (repoClient instanceof RepositoryApi) return repoClient;
  const _configuration = getApiConfig({ token, ...configuration, basePath });
  return new RepositoryApi(_configuration, _configuration.basePath, axios);
};

useRepoClient.propTypes = {
  token: PropTypes.string,
  basePath: PropTypes.string,
  repoClient: PropTypes.instanceOf(RepositoryApi),
  axios: PropTypes.instanceOf(AxiosInstance),
  /** *dcs-js* instance config */
  configuration: PropTypes.shape({
    apiKey: PropTypes.oneOfType([PropTypes.string, PropTypes.func, PropTypes.instanceOf(Promise)]),
    username: PropTypes.string,
    password: PropTypes.string,
    accessToken: PropTypes.oneOfType([PropTypes.string, PropTypes.func, PropTypes.instanceOf(Promise)]),
    basePath: PropTypes.string,
    baseOptions: PropTypes.object,
  })
};
```