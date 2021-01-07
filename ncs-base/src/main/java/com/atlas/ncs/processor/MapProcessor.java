package com.atlas.ncs.processor;

import java.util.Collection;
import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.stream.Collectors;

import com.app.rest.util.RestResponseUtil;
import com.atlas.mrg.constant.RestConstants;
import com.atlas.mrg.rest.attribute.MapCharacterAttributes;
import com.atlas.shared.rest.UriBuilder;

import rest.DataBody;
import rest.DataContainer;

public final class MapProcessor {
   public static CompletableFuture<Integer> countCharactersInMap(int worldId, int channelId, int mapId) {
      return UriBuilder.service(RestConstants.SERVICE)
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

   public static CompletableFuture<List<Integer>> getCharacterIdsInMap(int worldId, int channelId, int mapId) {
      return UriBuilder.service(RestConstants.SERVICE)
            .pathParam("worlds", worldId)
            .pathParam("channels", channelId)
            .pathParam("maps", mapId)
            .path("characters")
            .getAsyncRestClient(MapCharacterAttributes.class)
            .get()
            .thenApply(RestResponseUtil::result)
            .thenApply(DataContainer::dataList)
            .thenApply(Collection::stream)
            .thenApply(stream -> stream
                  .map(DataBody::getId)
                  .map(Integer::parseInt)
                  .collect(Collectors.toList())
            );
   }
}
