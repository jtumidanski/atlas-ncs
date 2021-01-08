package com.atlas.ncs.processor;

import java.util.Arrays;

import com.atlas.cos.rest.attribute.CharacterAttributes;
import com.atlas.mis.attribute.PortalAttributes;
import com.atlas.ncs.model.Character;
import com.atlas.ncs.model.Portal;

import rest.DataBody;

public final class ModelFactory {
   public static Portal createPortal(DataBody<PortalAttributes> body) {
      return new Portal(Integer.parseInt(body.getId()),
            body.getAttributes().name(),
            body.getAttributes().target(),
            body.getAttributes().type(),
            body.getAttributes().x(),
            body.getAttributes().y(),
            body.getAttributes().targetMap(),
            body.getAttributes().scriptName()
      );
   }

   public static Character createCharacter(DataBody<CharacterAttributes> body) {
      return new Character(Integer.parseInt(body.getId()),
            body.getAttributes().name(),
            body.getAttributes().level(),
            body.getAttributes().mapId(),
            body.getAttributes().x(),
            body.getAttributes().y(),
            body.getAttributes().hp(),
            body.getAttributes().gender(),
            body.getAttributes().jobId(),
            body.getAttributes().face(),
            body.getAttributes().hair(),
            Arrays.stream(body.getAttributes().sp().split(",")).mapToInt(Integer::parseInt).sum(),
            body.getAttributes().meso());
   }
}
