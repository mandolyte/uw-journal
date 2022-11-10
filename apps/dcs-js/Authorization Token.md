Sandbox: https://codesandbox.io/s/dcs-js-classes-auth-playground-m141om

```js
import { UserApi } from "dcs-js";
import "./styles.css";

async function apiPlayground() {
  const basePath = "https://qa.door43.org/api/v1";
  const username = "";
  const password = "";
  const tokenName = "dcs-js-test";

  if (!username || !password) {
    console.error(
      "No username or password provided: Please provide a username and password first on lines 6 and 7."
    );
    alert("No username or password provided");
    return;
  }

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
  console.log("Token=", token)

```