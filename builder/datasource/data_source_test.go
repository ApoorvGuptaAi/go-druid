package datasource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadUnsupportedType(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte("{\"type\": \"blahblahType\"}"))

	assert.Nil(f, "filter should be nil")
	assert.NotNil(err, "error should not be nil")
	assert.Error(err, "unsupported datasource type")
}

func TestLoadQuery(t *testing.T) {
	assert := assert.New(t)

        var qry string

	qry = `{     "type": "query",     "query": {         "dataSource": {             "type": "table",             "name": "multi_tenant_join_pred_label_regression"         },         "dimensions": [{             "type": "default",             "dimension": "__truera_model_id__",             "outputType": "STRING"         }],         "intervals": {             "type": "intervals",             "intervals": [                 "2022-01-12T10:59:51Z/2024-06-29T10:59:51Z"             ]         },         "queryType": "groupBy",         "granularity": "day"     } }`
	f, err := Load([]byte(qry))

	assert.Nil(f, "filter should be nil")
	assert.Nil(err, "error should not be nil")
}
