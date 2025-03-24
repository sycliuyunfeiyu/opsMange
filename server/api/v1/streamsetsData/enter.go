package streamsetsData

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ StreamsetsDataApi }

var sdcService = service.ServiceGroupApp.StreamsetsDataServiceGroup.StreamsetsDataService
