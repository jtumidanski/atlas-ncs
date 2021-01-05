package com.atlas.ncs.processor;

import java.util.List;
import java.util.concurrent.CompletableFuture;

import com.app.rest.util.RestResponseUtil;
import com.atlas.mrg.rest.attribute.MapCharacterAttributes;
import com.atlas.shared.rest.RestService;
import com.atlas.shared.rest.UriBuilder;

import rest.DataContainer;

public final class MapProcessor {
   public static CompletableFuture<Integer> countCharactersInMap(int worldId, int channelId, int mapId) {
      return UriBuilder.service(RestService.MAP_REGISTRY)
            .pathParam("worlds", worldId)
            .pathParam("channels", channelId)
            .pathParam("maps", mapId)
            .path("characters")
            .getAsyncRestClient(MapCharacterAttributes.class)
            .get()
            .thenApply(RestResponseUtil::result)
            .thenApply(DataContainer::dataList)
            .thenApply(List::size);
   }
}
