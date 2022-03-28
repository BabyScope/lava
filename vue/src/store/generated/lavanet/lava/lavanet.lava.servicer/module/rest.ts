/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface ServicerBlockDeadlineForCallback {
  deadline?: ServicerBlockNum;
}

export interface ServicerBlockNum {
  /** @format uint64 */
  num?: string;
}

export interface ServicerCurrentSessionStart {
  block?: ServicerBlockNum;
}

export interface ServicerEarliestSessionStart {
  block?: ServicerBlockNum;
}

export type ServicerMsgProofOfWorkResponse = object;

export type ServicerMsgStakeServicerResponse = object;

export type ServicerMsgUnstakeServicerResponse = object;

/**
 * Params defines the parameters for the module.
 */
export interface ServicerParams {
  /** @format uint64 */
  minStake?: string;

  /** @format uint64 */
  coinsPerCU?: string;

  /** @format uint64 */
  unstakeHoldBlocks?: string;

  /** @format uint64 */
  fraudStakeSlashingFactor?: string;

  /** @format uint64 */
  fraudSlashingAmount?: string;

  /** @format uint64 */
  servicersToPairCount?: string;

  /** @format uint64 */
  sessionBlocks?: string;

  /** @format uint64 */
  sessionsToSave?: string;

  /** @format uint64 */
  sessionBlocksOverlap?: string;
}

export interface ServicerPreviousSessionBlocks {
  /** @format uint64 */
  blocksNum?: string;
  changeBlock?: ServicerBlockNum;

  /** @format uint64 */
  overlapBlocks?: string;
}

export interface ServicerQueryAllSessionPaymentsResponse {
  sessionPayments?: ServicerSessionPayments[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface ServicerQueryAllSessionStorageForSpecResponse {
  sessionStorageForSpec?: ServicerSessionStorageForSpec[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface ServicerQueryAllSessionStoragesForSpecResponse {
  storages?: ServicerSessionStorageForSpec[];
}

export interface ServicerQueryAllSpecStakeStorageResponse {
  specStakeStorage?: ServicerSpecStakeStorage[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface ServicerQueryAllStakeMapResponse {
  stakeMap?: ServicerStakeMap[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface ServicerQueryAllUniquePaymentStorageUserServicerResponse {
  uniquePaymentStorageUserServicer?: ServicerUniquePaymentStorageUserServicer[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface ServicerQueryAllUnstakingServicersAllSpecsResponse {
  UnstakingServicersAllSpecs?: ServicerUnstakingServicersAllSpecs[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface ServicerQueryAllUserPaymentStorageResponse {
  userPaymentStorage?: ServicerUserPaymentStorage[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface ServicerQueryGetBlockDeadlineForCallbackResponse {
  BlockDeadlineForCallback?: ServicerBlockDeadlineForCallback;
}

export interface ServicerQueryGetCurrentSessionStartResponse {
  CurrentSessionStart?: ServicerCurrentSessionStart;
}

export interface ServicerQueryGetEarliestSessionStartResponse {
  EarliestSessionStart?: ServicerEarliestSessionStart;
}

export interface ServicerQueryGetPairingResponse {
  servicers?: ServicerStakeStorage;
}

export interface ServicerQueryGetPreviousSessionBlocksResponse {
  PreviousSessionBlocks?: ServicerPreviousSessionBlocks;
}

export interface ServicerQueryGetSessionPaymentsResponse {
  sessionPayments?: ServicerSessionPayments;
}

export interface ServicerQueryGetSessionStorageForSpecResponse {
  sessionStorageForSpec?: ServicerSessionStorageForSpec;
}

export interface ServicerQueryGetSpecStakeStorageResponse {
  specStakeStorage?: ServicerSpecStakeStorage;
}

export interface ServicerQueryGetStakeMapResponse {
  stakeMap?: ServicerStakeMap;
}

export interface ServicerQueryGetUniquePaymentStorageUserServicerResponse {
  uniquePaymentStorageUserServicer?: ServicerUniquePaymentStorageUserServicer;
}

export interface ServicerQueryGetUnstakingServicersAllSpecsResponse {
  UnstakingServicersAllSpecs?: ServicerUnstakingServicersAllSpecs;
}

export interface ServicerQueryGetUserPaymentStorageResponse {
  userPaymentStorage?: ServicerUserPaymentStorage;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface ServicerQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: ServicerParams;
}

export interface ServicerQuerySessionStorageForAllSpecsResponse {
  servicers?: ServicerStakeStorage;
}

export interface ServicerQueryStakedServicersResponse {
  stakeStorage?: ServicerStakeStorage;
  output?: string;
}

export interface ServicerQueryVerifyPairingResponse {
  valid?: boolean;
  overlap?: boolean;
}

export interface ServicerRelayReply {
  /** @format byte */
  data?: string;

  /** @format byte */
  sig?: string;
}

export interface ServicerRelayRequest {
  /** @format int64 */
  specId?: number;

  /** @format int64 */
  apiId?: number;

  /** @format uint64 */
  sessionId?: string;

  /** @format uint64 */
  cuSum?: string;

  /** @format byte */
  data?: string;

  /** @format byte */
  sig?: string;
  servicer?: string;

  /** @format int64 */
  blockHeight?: string;
}

export interface ServicerSessionPayments {
  index?: string;
  usersPayments?: ServicerUserPaymentStorage[];
}

export interface ServicerSessionStorageForSpec {
  index?: string;
  stakeStorage?: ServicerStakeStorage;
}

export interface ServicerSpecName {
  name?: string;
}

export interface ServicerSpecStakeStorage {
  index?: string;
  stakeStorage?: ServicerStakeStorage;
}

export interface ServicerStakeMap {
  index?: string;

  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  stake?: V1Beta1Coin;
  deadline?: ServicerBlockNum;
  operatorAddresses?: string[];
}

export interface ServicerStakeStorage {
  staked?: ServicerStakeMap[];
}

export interface ServicerUniquePaymentStorageUserServicer {
  index?: string;

  /** @format uint64 */
  block?: string;
}

export interface ServicerUnstakingServicersAllSpecs {
  /** @format uint64 */
  id?: string;
  unstaking?: ServicerStakeMap;
  specStakeStorage?: ServicerSpecStakeStorage;
}

export interface ServicerUserPaymentStorage {
  index?: string;
  uniquePaymentStorageUserServicer?: ServicerUniquePaymentStorageUserServicer[];

  /** @format uint64 */
  totalCU?: string;

  /** @format uint64 */
  session?: string;
}

/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
  denom?: string;
  amount?: string;
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  countTotal?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /** @format byte */
  nextKey?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title servicer/block_deadline_for_callback.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryAllSessionStoragesForSpec
   * @summary Queries a list of AllSessionStoragesForSpec items.
   * @request GET:/lavanet/lava/servicer/all_session_storages_for_spec/{specName}
   */
  queryAllSessionStoragesForSpec = (specName: string, params: RequestParams = {}) =>
    this.request<ServicerQueryAllSessionStoragesForSpecResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/all_session_storages_for_spec/${specName}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryBlockDeadlineForCallback
   * @summary Queries a BlockDeadlineForCallback by index.
   * @request GET:/lavanet/lava/servicer/block_deadline_for_callback
   */
  queryBlockDeadlineForCallback = (params: RequestParams = {}) =>
    this.request<ServicerQueryGetBlockDeadlineForCallbackResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/block_deadline_for_callback`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCurrentSessionStart
   * @summary Queries a CurrentSessionStart by index.
   * @request GET:/lavanet/lava/servicer/current_session_start
   */
  queryCurrentSessionStart = (params: RequestParams = {}) =>
    this.request<ServicerQueryGetCurrentSessionStartResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/current_session_start`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEarliestSessionStart
   * @summary Queries a EarliestSessionStart by index.
   * @request GET:/lavanet/lava/servicer/earliest_session_start
   */
  queryEarliestSessionStart = (params: RequestParams = {}) =>
    this.request<ServicerQueryGetEarliestSessionStartResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/earliest_session_start`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetPairing
   * @summary Queries a list of GetPairing items.
   * @request GET:/lavanet/lava/servicer/get_pairing/{specName}/{userAddr}
   */
  queryGetPairing = (specName: string, userAddr: string, params: RequestParams = {}) =>
    this.request<ServicerQueryGetPairingResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/get_pairing/${specName}/${userAddr}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/lavanet/lava/servicer/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<ServicerQueryParamsResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPreviousSessionBlocks
   * @summary Queries a PreviousSessionBlocks by index.
   * @request GET:/lavanet/lava/servicer/previous_session_blocks
   */
  queryPreviousSessionBlocks = (params: RequestParams = {}) =>
    this.request<ServicerQueryGetPreviousSessionBlocksResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/previous_session_blocks`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySessionPaymentsAll
   * @summary Queries a list of SessionPayments items.
   * @request GET:/lavanet/lava/servicer/session_payments
   */
  querySessionPaymentsAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ServicerQueryAllSessionPaymentsResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/session_payments`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySessionPayments
   * @summary Queries a SessionPayments by index.
   * @request GET:/lavanet/lava/servicer/session_payments/{index}
   */
  querySessionPayments = (index: string, params: RequestParams = {}) =>
    this.request<ServicerQueryGetSessionPaymentsResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/session_payments/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySessionStorageForAllSpecs
   * @summary Queries a list of SessionStorageForAllSpecs items.
   * @request GET:/lavanet/lava/servicer/session_storage_for_all_specs/{blockNum}
   */
  querySessionStorageForAllSpecs = (blockNum: string, params: RequestParams = {}) =>
    this.request<ServicerQuerySessionStorageForAllSpecsResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/session_storage_for_all_specs/${blockNum}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySessionStorageForSpecAll
   * @summary Queries a list of SessionStorageForSpec items.
   * @request GET:/lavanet/lava/servicer/session_storage_for_spec
   */
  querySessionStorageForSpecAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ServicerQueryAllSessionStorageForSpecResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/session_storage_for_spec`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySessionStorageForSpec
   * @summary Queries a SessionStorageForSpec by index.
   * @request GET:/lavanet/lava/servicer/session_storage_for_spec/{index}
   */
  querySessionStorageForSpec = (index: string, params: RequestParams = {}) =>
    this.request<ServicerQueryGetSessionStorageForSpecResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/session_storage_for_spec/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySpecStakeStorageAll
   * @summary Queries a list of SpecStakeStorage items.
   * @request GET:/lavanet/lava/servicer/spec_stake_storage
   */
  querySpecStakeStorageAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ServicerQueryAllSpecStakeStorageResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/spec_stake_storage`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySpecStakeStorage
   * @summary Queries a SpecStakeStorage by index.
   * @request GET:/lavanet/lava/servicer/spec_stake_storage/{index}
   */
  querySpecStakeStorage = (index: string, params: RequestParams = {}) =>
    this.request<ServicerQueryGetSpecStakeStorageResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/spec_stake_storage/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryStakeMapAll
   * @summary Queries a list of StakeMap items.
   * @request GET:/lavanet/lava/servicer/stake_map
   */
  queryStakeMapAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ServicerQueryAllStakeMapResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/stake_map`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryStakeMap
   * @summary Queries a StakeMap by index.
   * @request GET:/lavanet/lava/servicer/stake_map/{index}
   */
  queryStakeMap = (index: string, params: RequestParams = {}) =>
    this.request<ServicerQueryGetStakeMapResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/stake_map/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryStakedServicers
   * @summary Queries a list of StakedServicers items.
   * @request GET:/lavanet/lava/servicer/staked_servicers/{specName}
   */
  queryStakedServicers = (specName: string, params: RequestParams = {}) =>
    this.request<ServicerQueryStakedServicersResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/staked_servicers/${specName}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUniquePaymentStorageUserServicerAll
   * @summary Queries a list of UniquePaymentStorageUserServicer items.
   * @request GET:/lavanet/lava/servicer/unique_payment_storage_user_servicer
   */
  queryUniquePaymentStorageUserServicerAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ServicerQueryAllUniquePaymentStorageUserServicerResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/unique_payment_storage_user_servicer`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUniquePaymentStorageUserServicer
   * @summary Queries a UniquePaymentStorageUserServicer by index.
   * @request GET:/lavanet/lava/servicer/unique_payment_storage_user_servicer/{index}
   */
  queryUniquePaymentStorageUserServicer = (index: string, params: RequestParams = {}) =>
    this.request<ServicerQueryGetUniquePaymentStorageUserServicerResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/unique_payment_storage_user_servicer/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUnstakingServicersAllSpecsAll
   * @summary Queries a list of UnstakingServicersAllSpecs items.
   * @request GET:/lavanet/lava/servicer/unstaking_servicers_all_specs
   */
  queryUnstakingServicersAllSpecsAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ServicerQueryAllUnstakingServicersAllSpecsResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/unstaking_servicers_all_specs`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUnstakingServicersAllSpecs
   * @summary Queries a UnstakingServicersAllSpecs by id.
   * @request GET:/lavanet/lava/servicer/unstaking_servicers_all_specs/{id}
   */
  queryUnstakingServicersAllSpecs = (id: string, params: RequestParams = {}) =>
    this.request<ServicerQueryGetUnstakingServicersAllSpecsResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/unstaking_servicers_all_specs/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserPaymentStorageAll
   * @summary Queries a list of UserPaymentStorage items.
   * @request GET:/lavanet/lava/servicer/user_payment_storage
   */
  queryUserPaymentStorageAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ServicerQueryAllUserPaymentStorageResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/user_payment_storage`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserPaymentStorage
   * @summary Queries a UserPaymentStorage by index.
   * @request GET:/lavanet/lava/servicer/user_payment_storage/{index}
   */
  queryUserPaymentStorage = (index: string, params: RequestParams = {}) =>
    this.request<ServicerQueryGetUserPaymentStorageResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/user_payment_storage/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryVerifyPairing
   * @summary Queries a list of VerifyPairing items.
   * @request GET:/lavanet/lava/servicer/verify_pairing/{spec}/{userAddr}/{servicerAddr}/{blockNum}
   */
  queryVerifyPairing = (
    spec: string,
    userAddr: string,
    servicerAddr: string,
    blockNum: string,
    params: RequestParams = {},
  ) =>
    this.request<ServicerQueryVerifyPairingResponse, RpcStatus>({
      path: `/lavanet/lava/servicer/verify_pairing/${spec}/${userAddr}/${servicerAddr}/${blockNum}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
