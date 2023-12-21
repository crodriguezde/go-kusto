package query

import (
	"time"

	"github.com/crodriguezde/go-kusto/pkg/value"
)

const ResultsProgressiveEnabledValue = "results_progressive_enabled"
const NoRequestTimeoutValue = "norequesttimeout"
const NoTruncationValue = "notruncation"
const ServerTimeoutValue = "servertimeout"
const DeferPartialQueryFailuresValue = "deferpartialqueryfailures"
const MaxMemoryConsumptionPerQueryPerNodeValue = "max_memory_consumption_per_query_per_node"
const MaxMemoryConsumptionPerIteratorValue = "maxmemoryconsumptionperiterator"
const MaxOutputColumnsValue = "maxoutputcolumns"
const PushSelectionThroughAggregationValue = "push_selection_through_aggregation"
const QueryCursorAfterDefaultValue = "query_cursor_after_default"
const QueryCursorBeforeOrAtDefaultValue = "query_cursor_before_or_at_default"
const QueryCursorCurrentValue = "query_cursor_current"
const QueryCursorDisabledValue = "query_cursor_disabled"
const QueryCursorScopedTablesValue = "query_cursor_scoped_tables"
const QueryDatascopeValue = "query_datascope"
const QueryDateTimeScopeColumnValue = "query_datetimescope_column"
const QueryDateTimeScopeFromValue = "query_datetimescope_from"
const QueryDateTimeScopeToValue = "query_datetimescope_to"
const ClientMaxRedirectCountValue = "client_max_redirect_count"
const MaterializedViewShuffleValue = "materialized_view_shuffle"
const QueryBinAutoAtValue = "query_bin_auto_at"
const QueryBinAutoSizeValue = "query_bin_auto_size"
const QueryDistributionNodesSpanValue = "query_distribution_nodes_span"
const QueryFanoutNodesPercentValue = "query_fanout_nodes_percent"
const QueryFanoutThreadsPercentValue = "query_fanout_threads_percent"
const QueryForceRowLevelSecurityValue = "query_force_row_level_security"
const QueryLanguageValue = "query_language"
const QueryLogQueryParametersValue = "query_log_query_parameters"
const QueryMaxEntitiesInUnionValue = "query_max_entities_in_union"
const QueryNowValue = "query_now"
const QueryPythonDebugValue = "query_python_debug"
const QueryResultsApplyGetschemaValue = "query_results_apply_getschema"
const QueryResultsCacheMaxAgeValue = "query_results_cache_max_age"
const QueryResultsCachePerShardValue = "query_results_cache_per_shard"
const QueryResultsProgressiveRowCountValue = "query_results_progressive_row_count"
const QueryResultsProgressiveUpdatePeriodValue = "query_results_progressive_update_period"
const QueryTakeMaxRecordsValue = "query_take_max_records"
const QueryConsistencyValue = "queryconsistency"
const RequestAppNameValue = "request_app_name"
const RequestBlockRowLevelSecurityValue = "request_block_row_level_security"
const RequestCalloutDisabledValue = "request_callout_disabled"
const RequestDescriptionValue = "request_description"
const RequestExternalTableDisabledValue = "request_external_table_disabled"
const RequestImpersonationDisabledValue = "request_impersonation_disabled"
const RequestReadonlyValue = "request_readonly"
const RequestRemoteEntitiesDisabledValue = "request_remote_entities_disabled"
const RequestSandboxedExecutionDisabledValue = "request_sandboxed_execution_disabled"
const RequestUserValue = "request_user"
const TruncationMaxRecordsValue = "truncation_max_records"
const TruncationMaxSizeValue = "truncation_max_size"
const ValidatePermissionsValue = "validate_permissions"

type RequestProperties struct {
	Options         map[string]interface{}
	Parameters      map[string]string
	Application     string                 `json:"-"`
	User            string                 `json:"-"`
	QueryParameters map[string]value.Value `json:"-"`
	ClientRequestID string                 `json:"-"`
}

type QueryOptions struct {
	RequestProperties *RequestProperties
}

type QueryOption func(q *QueryOptions) error

// ClientRequestID sets the x-ms-client-request-id header, and can be used to identify the request in the `.show queries` output.
func ClientRequestID(clientRequestID string) QueryOption {
	return func(q *QueryOptions) error {
		q.RequestProperties.ClientRequestID = clientRequestID
		return nil
	}
}

// Application sets the x-ms-app header, and can be used to identify the application making the request in the `.show queries` output.
func Application(appName string) QueryOption {
	return func(q *QueryOptions) error {
		q.RequestProperties.Application = appName
		return nil
	}
}

// User sets the x-ms-user header, and can be used to identify the user making the request in the `.show queries` output.
func User(userName string) QueryOption {
	return func(q *QueryOptions) error {
		q.RequestProperties.User = userName
		return nil
	}
}

// NoRequestTimeout enables setting the request timeout to its maximum value.
func NoRequestTimeout() QueryOption {
	return func(q *QueryOptions) error {
		q.RequestProperties.Options[NoRequestTimeoutValue] = true
		return nil
	}
}

// NoTruncation enables suppressing truncation of the query results returned to the caller.
func NoTruncation() QueryOption {
	return func(q *QueryOptions) error {
		q.RequestProperties.Options[NoTruncationValue] = true
		return nil
	}
}

// ResultsProgressiveEnabled enables the progressive query stream.
func ResultsProgressiveEnabled() QueryOption {
	return func(q *QueryOptions) error {
		q.RequestProperties.Options[ResultsProgressiveEnabledValue] = true
		return nil
	}
}

// ServerTimeout overrides the default request timeout.
func ServerTimeout(d time.Duration) QueryOption {
	return func(q *QueryOptions) error {
		q.RequestProperties.Options[ServerTimeoutValue] = value.Timespan{Valid: true, Value: d}.Marshal()
		return nil
	}
}

// DeferPartialQueryFailures disables reporting partial query failures as part of the result set.
func DeferPartialQueryFailures() QueryOption {
	return func(q *QueryOptions) error {
		q.RequestProperties.Options[DeferPartialQueryFailuresValue] = true
		return nil
	}
}

// MaxMemoryConsumptionPerQueryPerNode overrides the default maximum amount of memory a whole query
// may allocate per node.
func MaxMemoryConsumptionPerQueryPerNode(i uint64) QueryOption {
	return func(q *QueryOptions) error {
		q.RequestProperties.Options[MaxMemoryConsumptionPerQueryPerNodeValue] = i
		return nil
	}
}

// MaxMemoryConsumptionPerIterator overrides the default maximum amount of memory a query operator may allocate.
func MaxMemoryConsumptionPerIterator(i uint64) QueryOption {
	return func(q *QueryOptions) error {
		q.RequestProperties.Options[MaxMemoryConsumptionPerIteratorValue] = i
		return nil
	}
}
