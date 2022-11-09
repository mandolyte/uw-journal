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

Example for auth:
https://codesandbox.io/s/dcs-js-classes-auth-playground-m141om?file=/src/index.js

Code in above:
```js
import { OrganizationApi, UserApi } from "dcs-js";
import "./styles.css";

async function apiPlayground() {
  const basePath = "https://qa.door43.org/api/v1";
  const username = "";
  const password = "";
  const tokenName = "dcs-js-test";
  const newOrgName = `dcs-js-test-org`;

  const userClient = new UserApi({ basePath, username, password });

  const deletePreviousTokens = async (tokenName) => {
    const previousTokens = await userClient
      .userGetTokens({ username })
      .then(({ data }) => data);

    if (!previousTokens?.length) return false;

    previousTokens.forEach((token) => {
      if (token.name === tokenName)
        userClient.userDeleteAccessToken({ username, token: token.id });
    });

    return true;
  };

  const getToken = async (tokenName) => {
    await deletePreviousTokens(tokenName);
    return await userClient
      .userCreateToken({
        username,
        userCreateToken: { name: tokenName }
      })
      .then(({ data }) => data);
  };

  const token = await getToken(tokenName);

  /* Instantiate the required DCS client. */
  const organizationClient = new OrganizationApi({
    /* Set the desired DCS server path */
    basePath,
    apiKey: (keyName) =>
      keyName === "Authorization" ? `token ${token.sha1}` : undefined
  });

  const createdOrg = await organizationClient
    .orgCreate({
      organization: { username: newOrgName }
    })
    .then(({ data }) => data);

  const deleteOrg = async (orgName) =>
    await organizationClient.orgDelete({ org: orgName });

  /* Clean up generated org */
  await deleteOrg(newOrgName);

  /* Use the returned data */
  const printArea = document.getElementById("print-area");
  const printableData = JSON.stringify(createdOrg, null, 4);
  printArea.innerHTML = `<pre>${printableData}</pre>`;
}

apiPlayground();

```