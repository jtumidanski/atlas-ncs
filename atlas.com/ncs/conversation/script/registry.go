package script

import (
	"errors"
	"sync"
)

type Registry struct {
	registry map[uint32]Script
}

var once sync.Once
var registry *Registry

func GetRegistry() *Registry {
	once.Do(func() {
		registry = initRegistry()
	})
	return registry
}

func initRegistry() *Registry {
	s := &Registry{make(map[uint32]Script)}
	s.addConversation(AFamiliarLady{})
	s.addConversation(APileOfBlueFlowers{})
	s.addConversation(APileOfFlowers{})
	s.addConversation(APileOfHerbs{})
	s.addConversation(APileOfPinkFlowers{})
	s.addConversation(APileOfWhiteFlowers{})
	s.addConversation(AbdullahVIII{})
	s.addConversation(AceOfHearts{})
	s.addConversation(Agatha{})
	s.addConversation(AlcandosCabinet{})
	s.addConversation(Alcaster{})
	s.addConversation(Ali{})
	s.addConversation(Amon{})
	s.addConversation(AmosTheStrong{})
	s.addConversation(AncientIcyStone{})
	s.addConversation(Andy{})
	s.addConversation(Ardin{})
	s.addConversation(Areda{})
	s.addConversation(AriantPrivateHouse1{})
	s.addConversation(AriantPrivateHouse1Cupboard{})
	s.addConversation(AriantPrivateHouse2{})
	s.addConversation(AriantPrivateHouse2Cupboard{})
	s.addConversation(AriantPrivateHouse4{})
	s.addConversation(AriantPrivateHouse4Cupboard{})
	s.addConversation(AriantPrivateHouse6{})
	s.addConversation(AriantPrivateHouse6Cupboard{})
	s.addConversation(Arturo{})
	s.addConversation(ArwenTheFairy{})
	s.addConversation(Asesson{})
	s.addConversation(Asia{})
	s.addConversation(AssistantBlue{})
	s.addConversation(AssistantBlue2{})
	s.addConversation(AssistantCheng{})
	s.addConversation(AssistantRed{})
	s.addConversation(AssistantRed2{})
	s.addConversation(AthenaPierce{})
	s.addConversation(AthenaPierceDemo{})
	s.addConversation(Aura{})
	s.addConversation(BabyMilkCow1{})
	s.addConversation(BabyMilkCow2{})
	s.addConversation(BabyMoonBunny{})
	s.addConversation(Bart{})
	s.addConversation(Bedin{})
	s.addConversation(BigHeadward{})
	s.addConversation(BlockedEntrance{})
	s.addConversation(BossKitty{})
	s.addConversation(BowmanJobInstructor{})
	s.addConversation(BowmanJobInstructorExit{})
	s.addConversation(BranchSnowman{})
	s.addConversation(Brittany{})
	s.addConversation(BulletinBoard{})
	s.addConversation(Bush1{})
	s.addConversation(Bush2{})
	s.addConversation(Bush3{})
	s.addConversation(Bush4{})
	s.addConversation(Bush5{})
	s.addConversation(Byron{})
	s.addConversation(CamelCab{})
	s.addConversation(Carson{})
	s.addConversation(Carta{})
	s.addConversation(Casey{})
	s.addConversation(CenterOfTheMagicPentagram{})
	s.addConversation(Cesar3{})
	s.addConversation(Chef{})
	s.addConversation(Cherry{})
	s.addConversation(ChiefTatamo{})
	s.addConversation(Chris{})
	s.addConversation(ChunJi{})
	s.addConversation(Cliff{})
	s.addConversation(Cloy{})
	s.addConversation(Coco{})
	s.addConversation(ControlDevice{})
	s.addConversation(Corba{})
	s.addConversation(CornerOfTheMagicLibrary{})
	s.addConversation(CornerOfTheMagicLibrary2{})
	s.addConversation(Crane{})
	s.addConversation(CrawlsWithBalrog{})
	s.addConversation(CrumblingStatue{})
	s.addConversation(Cygnus{})
	s.addConversation(DancesWithBalrog{})
	s.addConversation(DancesWithBalrogDemo{})
	s.addConversation(DangerZoneTaxi{})
	s.addConversation(DarkLordDemo{})
	s.addConversation(Desk{})
	s.addConversation(DimensionalMirror{})
	s.addConversation(DocumentRoll{})
	s.addConversation(Dolphin{})
	s.addConversation(Dolphin2{})
	s.addConversation(Dunamis{})
	s.addConversation(Duru{})
	s.addConversation(Eckhart{})
	s.addConversation(EckhartStatue{})
	s.addConversation(Egnet{})
	s.addConversation(Eleanor{})
	s.addConversation(Eleska{})
	s.addConversation(EllinForestMilepost{})
	s.addConversation(EncryptedSlateOfTheSquad{})
	s.addConversation(EntranceOfSealedShrine{})
	s.addConversation(Erin{})
	s.addConversation(FallenKnight{})
	s.addConversation(FirstEOSRock{})
	s.addConversation(FirstMagicPentagram{})
	s.addConversation(FirstPipeHandle{})
	s.addConversation(FountainOfLife{})
	s.addConversation(FourthEOSRock{})
	s.addConversation(Francis{})
	s.addConversation(Francois{})
	s.addConversation(Francis2{})
	s.addConversation(FranzTheOwner{})
	s.addConversation(Gaga{})
	s.addConversation(Geras{})
	s.addConversation(GhostHunterBob{})
	s.addConversation(Gina{})
	s.addConversation(GrandpaMoonBunny{})
	s.addConversation(GrendelTheReallyOld{})
	s.addConversation(GrendelTheReallyOldDemo{})
	s.addConversation(Gritto{})
	s.addConversation(HanTheBroker{})
	s.addConversation(Harmonia{})
	s.addConversation(HarpStringA{})
	s.addConversation(HarpStringB{})
	s.addConversation(HarpStringC{})
	s.addConversation(HarpStringD{})
	s.addConversation(HarpStringE{})
	s.addConversation(HarpStringF{})
	s.addConversation(HarpStringG{})
	s.addConversation(Harry{})
	s.addConversation(HarryCoconut{})
	s.addConversation(Hawkeye{})
	s.addConversation(HawkeyeStatue{})
	s.addConversation(Heena{})
	s.addConversation(Hellin{})
	s.addConversation(HenesysForest{})
	s.addConversation(HenesysForest2{})
	s.addConversation(Heracle{})
	s.addConversation(HiddenNote{})
	s.addConversation(HiddenNote2{})
	s.addConversation(Hikari{})
	s.addConversation(HolyStone{})
	s.addConversation(HotelReceptionist{})
	s.addConversation(HumanoidA{})
	s.addConversation(IncompleteMagicSquare{})
	s.addConversation(InsignificantBeing{})
	s.addConversation(Irena{})
	s.addConversation(IrenaStatue{})
	s.addConversation(IsaTheStationGuide{})
	s.addConversation(Jack{})
	s.addConversation(Jake{})
	s.addConversation(Jane{})
	s.addConversation(Jano{})
	s.addConversation(Jean{})
	s.addConversation(Jeff{})
	s.addConversation(Jiyur{})
	s.addConversation(JMFromThaStreetz{})
	s.addConversation(Joel{})
	s.addConversation(Kanderune{})
	s.addConversation(Karcasa{})
	s.addConversation(Karen{})
	s.addConversation(Keeny{})
	s.addConversation(Kenta{})
	s.addConversation(KingPepe{})
	s.addConversation(KinoKonoko{})
	s.addConversation(Kiridu{})
	s.addConversation(Kiriko{})
	s.addConversation(KiriruEreve{})
	s.addConversation(KiruEreve{})
	s.addConversation(KiruToOrbis{})
	s.addConversation(KiruOrbis{})
	s.addConversation(Kiruru{})
	s.addConversation(KnightArmor{})
	s.addConversation(Konpei{})
	s.addConversation(KonpeiExit{})
	s.addConversation(KonpeiExpeditionSuccessExit{})
	s.addConversation(KiriruEllinia{})
	s.addConversation(KyrinDemo{})
	s.addConversation(LePetitPrince{})
	s.addConversation(Lea{})
	s.addConversation(Legor{})
	s.addConversation(Lenario{})
	s.addConversation(Lila{})
	s.addConversation(Lira{})
	s.addConversation(Lisa{})
	s.addConversation(Loha{})
	s.addConversation(Lohd{})
	s.addConversation(LordJonathan{})
	s.addConversation(Louis{})
	s.addConversation(Luke{})
	s.addConversation(MachineApparatus{})
	s.addConversation(Maed{})
	s.addConversation(MagicianJobInstructor{})
	s.addConversation(MagicianJobInstructorExit{})
	s.addConversation(Maid{})
	s.addConversation(MapleLeafMarble{})
	s.addConversation(MarTheFairy{})
	s.addConversation(MarkTheToySoldier{})
	s.addConversation(Mel{})
	s.addConversation(MetalBucketSnowman{})
	s.addConversation(Meteorite1{})
	s.addConversation(Meteorite2{})
	s.addConversation(Meteorite3{})
	s.addConversation(Meteorite4{})
	s.addConversation(Meteorite5{})
	s.addConversation(Meteorite6{})
	s.addConversation(Mia{})
	s.addConversation(Mia2{})
	s.addConversation(Mihile{})
	s.addConversation(MihileStatue{})
	s.addConversation(Mimo{})
	s.addConversation(MinoTheOwner{})
	s.addConversation(MongFromKong{})
	s.addConversation(MonstrousLookingStatue{})
	s.addConversation(Moose{})
	s.addConversation(MooseExit{})
	s.addConversation(MotherMilkCow1{})
	s.addConversation(MotherMilkCow2{})
	s.addConversation(MrGoldstein{})
	s.addConversation(MrPickall{})
	s.addConversation(MrSmith{})
	s.addConversation(MrThunder{})
	s.addConversation(MsTan{})
	s.addConversation(MuLungDojoBulletinBoard{})
	s.addConversation(Mue{})
	s.addConversation(Muirhat{})
	s.addConversation(MysteriousStatue{})
	s.addConversation(Nara{})
	s.addConversation(Naran{})
	s.addConversation(Natalie{})
	s.addConversation(NautilusMidSizedTaxi{})
	s.addConversation(Neinheart{})
	s.addConversation(NeinheartDemo{})
	s.addConversation(Nella{})
	s.addConversation(Neru{})
	s.addConversation(Nuris{})
	s.addConversation(OlsonTheToySoldier{})
	s.addConversation(OrbisMagicSpot1{})
	s.addConversation(OrbisMagicSpot20{})
	s.addConversation(Oz{})
	s.addConversation(OzStatue{})
	s.addConversation(PalaceOasis{})
	s.addConversation(Papulatus{})
	s.addConversation(Parwen{})
	s.addConversation(Pason{})
	s.addConversation(Paul{})
	s.addConversation(Pelace{})
	s.addConversation(Perzen{})
	s.addConversation(Peter{})
	s.addConversation(Phil{})
	s.addConversation(Phyllia{})
	s.addConversation(PictureFrame{})
	s.addConversation(Pietra{})
	s.addConversation(Pietro{})
	s.addConversation(PirateJobInstructorExit{})
	s.addConversation(Pison{})
	s.addConversation(PowerBForeBowman{})
	s.addConversation(PowerBForeEllinia{})
	s.addConversation(PowerBForeHenesys{})
	s.addConversation(PowerBForeKerningCity{})
	s.addConversation(PowerBForeMagician{})
	s.addConversation(PowerBForeNautilus{})
	s.addConversation(PowerBForePerion{})
	s.addConversation(PowerBForePirate{})
	s.addConversation(PowerBForeThief{})
	s.addConversation(PowerBForeWarrior{})
	s.addConversation(PracticeChart{})
	s.addConversation(PrinceGiuseppe{})
	s.addConversation(Purin{})
	s.addConversation(PuroLithHarbor{})
	s.addConversation(PuroRien{})
	s.addConversation(PuroToLithHarbor{})
	s.addConversation(PuroToRien{})
	s.addConversation(QueensCabinet{})
	s.addConversation(Rain{})
	s.addConversation(Ramini{})
	s.addConversation(Refugee1{})
	s.addConversation(Refugee2{})
	s.addConversation(Refugee3{})
	s.addConversation(Refugee4{})
	s.addConversation(RegularCabEllinia{})
	s.addConversation(RegularCabHenesys{})
	s.addConversation(RegularCabKerningCity{})
	s.addConversation(RegularCabLithHarbor{})
	s.addConversation(RegularCabPerion{})
	s.addConversation(ReturningRock{})
	s.addConversation(Ria{})
	s.addConversation(Rini{})
	s.addConversation(Rius{})
	s.addConversation(RobertHolly{})
	s.addConversation(Robin{})
	s.addConversation(RobinTheHuntress{})
	s.addConversation(RollyExit{})
	s.addConversation(Romi{})
	s.addConversation(Rooney{})
	s.addConversation(Rosey{})
	s.addConversation(Rupi{})
	s.addConversation(Russellon{})
	s.addConversation(RussellonsDesk{})
	s.addConversation(Samuel{})
	s.addConversation(ScarfSnowman{})
	s.addConversation(Schegerazade{})
	s.addConversation(SeargantAnderson{})
	s.addConversation(SecondEOSRock{})
	s.addConversation(SecondMagicPentagram{})
	s.addConversation(SecondPipeHandle{})
	s.addConversation(SecretWall{})
	s.addConversation(Sejan{})
	s.addConversation(Sera{})
	s.addConversation(Shadrion{})
	s.addConversation(ShamanRock1{})
	s.addConversation(ShamanRock2{})
	s.addConversation(Shane{})
	s.addConversation(Shanks{})
	s.addConversation(SharenIIIsWill{})
	s.addConversation(Shawn{})
	s.addConversation(Shinsoo{})
	s.addConversation(ShinyStone{})
	s.addConversation(Shulynch{})
	s.addConversation(Shulynch2{})
	s.addConversation(Shuri{})
	s.addConversation(Simon{})
	s.addConversation(Sion{})
	s.addConversation(Sirin{})
	s.addConversation(Slyn{})
	s.addConversation(SmallStreetLight{})
	s.addConversation(SmallTreeStump{})
	s.addConversation(SnowmanExit{})
	s.addConversation(SparklingCrystal{})
	s.addConversation(Spinel{})
	s.addConversation(Spiruna{})
	s.addConversation(StaffSeargantCharlie{})
	s.addConversation(StrangeLookingStatue{})
	s.addConversation(StrawHatSnowman{})
	s.addConversation(SubwayExit{})
	s.addConversation(SubwayTrashCan1{})
	s.addConversation(SubwayTrashCan2{})
	s.addConversation(SubwayTrashCan3{})
	s.addConversation(SubwayTrashCan4{})
	s.addConversation(Sunny{})
	s.addConversation(Syras{})
	s.addConversation(Tangyoon{})
	s.addConversation(Tess{})
	s.addConversation(TheTicketGate{})
	s.addConversation(ThiefJobInstructor{})
	s.addConversation(ThiefJobInstructorExit{})
	s.addConversation(ThirdEOSRock{})
	s.addConversation(ThirdMagicPentagram{})
	s.addConversation(ThirdPipeHandle{})
	s.addConversation(ThomasSwift{})
	s.addConversation(ThreeRefugees{})
	s.addConversation(Tia{})
	s.addConversation(Tian{})
	s.addConversation(Tigun{})
	s.addConversation(TimeGate{})
	s.addConversation(Tombstone{})
	s.addConversation(Tommie{})
	s.addConversation(Tony{})
	s.addConversation(TrainerBartos{})
	s.addConversation(TrainerFrod{})
	s.addConversation(TrashCan{})
	s.addConversation(TreasureChestB1{})
	s.addConversation(TreasureChestB2{})
	s.addConversation(TreasureChestB3{})
	s.addConversation(TutorialLilin{})
	s.addConversation(TylusComplete{})
	s.addConversation(Vicious{})
	s.addConversation(Victoria{})
	s.addConversation(Vikin{})
	s.addConversation(VIPCabEllinia{})
	s.addConversation(VIPCabLithHarbor{})
	s.addConversation(WarriorJobInstructor{})
	s.addConversation(WarriorJobInstructorExit{})
	s.addConversation(WaterFilter{})
	s.addConversation(Weaver{})
	s.addConversation(Wisp{})
	s.addConversation(WolfGuard{})
	s.addConversation(WolfSpiritRyko{})
	s.addConversation(YokoYoko{})
	s.addConversation(YoungAthenaPierce{})
	s.addConversation(YoungAthenaPierce2{})
	s.addConversation(YuleteDefeated{})
	s.addConversation(Yuris{})
	return s
}

func (s *Registry) GetScript(npcId uint32) (*Script, error) {
	if val, ok := s.registry[npcId]; ok {
		return &val, nil
	}
	return nil, errors.New("unable to locate script")
}

func (s *Registry) addConversation(handler Script) {
	s.registry[handler.NPCId()] = handler
}
