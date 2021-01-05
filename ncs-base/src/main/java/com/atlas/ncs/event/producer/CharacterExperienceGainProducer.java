package com.atlas.ncs.event.producer;

import com.atlas.cos.constant.EventConstants;
import com.atlas.cos.event.CharacterExperienceEvent;
import com.atlas.ncs.EventProducerRegistry;

public final class CharacterExperienceGainProducer {
   private CharacterExperienceGainProducer() {
   }

   public static void gainExperience(int characterId, int gain) {
      EventProducerRegistry.getInstance()
            .send(CharacterExperienceEvent.class, EventConstants.TOPIC_CHARACTER_EXPERIENCE_EVENT, characterId,
                  new CharacterExperienceEvent(characterId, gain, 0, true, false, true));
   }
}
