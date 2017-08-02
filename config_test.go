package biigo

import "testing"

import "github.com/stretchr/testify/suite"

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(configTestSuite))
}

type configTestSuite struct {
	suite.Suite
	config Config
}

func (suite *configTestSuite) SetupTest() {
	app := &App{}
	app.LoadConfig("./test_configs")
	suite.config = app.Config()
}

func (suite *configTestSuite) TestString() {
	assert := suite.Assert()

	assert.Equal("value1", suite.config.String("field1", ""))
	assert.Equal("value2-2", suite.config.String("field2", ""))
}

func (suite *configTestSuite) TestJson() {
	js := struct {
		Field3 string `json:"field3"`
	}{}

	assert := suite.Assert()
	assert.Empty(suite.config.JSONUnmarshal("json", &js))
	assert.Equal("value3", js.Field3)
}
