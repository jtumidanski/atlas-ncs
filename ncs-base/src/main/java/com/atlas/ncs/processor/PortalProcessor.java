package com.atlas.ncs.processor;

import java.util.Collection;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.CompletableFuture;
import java.util.stream.Collectors;

import com.app.rest.util.RestResponseUtil;
import com.atlas.mis.attribute.PortalAttributes;
import com.atlas.ncs.model.Portal;
import com.atlas.shared.rest.RestService;
import com.atlas.shared.rest.UriBuilder;

import rest.DataContainer;

public final class PortalProcessor {
   private PortalProcessor() {
   }

   public static CompletableFuture<List<Portal>> getMapPortals(int mapId) {
      return UriBuilder.service(RestService.MAP_INFORMATION)
            .pathParam("maps", mapId)
            .path("portals")
            .getAsyncRestClient(PortalAttributes.class)
            .get()
            .thenApply(RestResponseUtil::result)
            .thenApply(DataContainer::dataList)
            .thenApply(Collection::stream)
            .thenApply(stream -> stream.map(ModelFactory::createPortal))
            .thenApply(stream -> stream.collect(Collectors.toList()));
   }

   public static CompletableFuture<Portal> getMapPortalByName(int mapId, String name) {
      return UriBuilder.service(RestService.MAP_INFORMATION)
            .pathParam("maps", mapId)
            .path("portals")
            .queryParam("name", name)
            .getAsyncRestClient(PortalAttributes.class)
            .get()
            .thenApply(RestResponseUtil::result)
            .thenApply(DataContainer::data)
            .thenApply(Optional::get)
            .thenApply(ModelFactory::createPortal);
   }

   public static CompletableFuture<Portal> getMapPortalById(int mapId, int portalId) {
      return UriBuilder.service(RestService.MAP_INFORMATION)
            .pathParam("maps", mapId)
            .pathParam("portals", portalId)
            .getAsyncRestClient(PortalAttributes.class)
            .get()
            .thenApply(RestResponseUtil::result)
            .thenApply(DataContainer::data)
            .thenApply(Optional::get)
            .thenApply(ModelFactory::createPortal);
   }

   protected static boolean isSpawnPoint(Portal portal) {
      return (portal.type() == 0 || portal.type() == 1) && portal.targetMap() == 999999999;
   }
}
