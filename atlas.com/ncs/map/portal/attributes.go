package portal

import "atlas-ncs/rest/response"

type dataContainer struct {
   data     response.DataSegment
   included response.DataSegment
}

type dataBody struct {
   Id         string           `json:"id"`
   Type       string     `json:"type"`
   Attributes attributes `json:"attributes"`
}

type attributes struct {
   Name       string `json:"name"`
   Target     string `json:"target"`
   Type       uint32 `json:"type"`
   X          int32 `json:"x"`
   Y          int32 `json:"y"`
   TargetMap  uint32 `json:"targetMap"`
   ScriptName string `json:"scriptName"`
}

func (a *dataContainer) UnmarshalJSON(data []byte) error {
   d, i, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyPortalData))
   if err != nil {
      return err
   }

   a.data = d
   a.included = i
   return nil
}

func (a *dataContainer) Data() *dataBody {
   if len(a.data) >= 1 {
      return a.data[0].(*dataBody)
   }
   return nil
}

func (a *dataContainer) DataList() []*dataBody {
   var r = make([]*dataBody, 0)
   for _, x := range a.data {
      r = append(r, x.(*dataBody))
   }
   return r
}

func EmptyPortalData() interface{} {
   return &dataBody{}
}
