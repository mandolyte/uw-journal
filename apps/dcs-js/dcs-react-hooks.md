# dcs-react-hooks

## misc notes

In dcs-js, possible org methods:
```ts
orgIsMember: async (org: string, username: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {


			
orgListCurrentUserOrgs: async (page?: number, limit?: number, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {


async orgListCurrentUserOrgs(page?: number, limit?: number, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => Promise<AxiosResponse<Array<Organization>>>> {
  const localVarAxiosArgs = await 
	OrganizationApiAxiosParamCreator(configuration)
      .orgListCurrentUserOrgs(page, limit, options);
  return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs :AxiosRequestConfig = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },


```