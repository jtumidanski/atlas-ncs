package registry

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/conversation/script/discrete"
	"errors"
	"sync"
)

type Registry struct {
	registry map[uint32]script.Script
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
	s := &Registry{make(map[uint32]script.Script)}
	s.addConversation(discrete.AFamiliarLady{})
	s.addConversation(discrete.APileOfBlueFlowers{})
	s.addConversation(discrete.APileOfFlowers{})
	s.addConversation(discrete.APileOfHerbs{})
	s.addConversation(discrete.APileOfPinkFlowers{})
	s.addConversation(discrete.APileOfWhiteFlowers{})
	s.addConversation(discrete.AbdullahVIII{})
	s.addConversation(discrete.AceOfHearts{})
	s.addConversation(discrete.Adonis{})
	s.addConversation(discrete.Agatha{})
	s.addConversation(discrete.AlcandosCabinet{})
	s.addConversation(discrete.Alcaster{})
	s.addConversation(discrete.Ali{})
	s.addConversation(discrete.Amon{})
	s.addConversation(discrete.AmosTheStrong{})
	s.addConversation(discrete.AncientIcyStone{})
	s.addConversation(discrete.Andy{})
	s.addConversation(discrete.Ardin{})
	s.addConversation(discrete.Areda{})
	s.addConversation(discrete.AriantPrivateHouse1{})
	s.addConversation(discrete.AriantPrivateHouse1Cupboard{})
	s.addConversation(discrete.AriantPrivateHouse2{})
	s.addConversation(discrete.AriantPrivateHouse2Cupboard{})
	s.addConversation(discrete.AriantPrivateHouse4{})
	s.addConversation(discrete.AriantPrivateHouse4Cupboard{})
	s.addConversation(discrete.AriantPrivateHouse6{})
	s.addConversation(discrete.AriantPrivateHouse6Cupboard{})
	s.addConversation(discrete.Arturo{})
	s.addConversation(discrete.ArwenTheFairy{})
	s.addConversation(discrete.Asesson{})
	s.addConversation(discrete.Asia{})
	s.addConversation(discrete.AssistantBlue{})
	s.addConversation(discrete.AssistantBlue2{})
	s.addConversation(discrete.AssistantCheng{})
	s.addConversation(discrete.AssistantRed{})
	s.addConversation(discrete.AssistantRed2{})
	s.addConversation(discrete.AstarothsDoorway{})
	s.addConversation(discrete.AthenaPierce{})
	s.addConversation(discrete.AthenaPierceDemo{})
	s.addConversation(discrete.Aura{})
	s.addConversation(discrete.BabyMilkCow1{})
	s.addConversation(discrete.BabyMilkCow2{})
	s.addConversation(discrete.BabyMoonBunny{})
	s.addConversation(discrete.Bart{})
	s.addConversation(discrete.Bedin{})
	s.addConversation(discrete.Bell{})
	s.addConversation(discrete.BigHeadward{})
	s.addConversation(discrete.BlockedEntrance{})
	s.addConversation(discrete.BossKitty{})
	s.addConversation(discrete.BowmanJobInstructor{})
	s.addConversation(discrete.BowmanJobInstructorExit{})
	s.addConversation(discrete.BowmanStatue{})
	s.addConversation(discrete.BranchSnowman{})
	s.addConversation(discrete.Brittany{})
	s.addConversation(discrete.BulletinBoard{})
	s.addConversation(discrete.Bush1{})
	s.addConversation(discrete.Bush2{})
	s.addConversation(discrete.Bush3{})
	s.addConversation(discrete.Bush4{})
	s.addConversation(discrete.Bush5{})
	s.addConversation(discrete.Byron{})
	s.addConversation(discrete.CamelCab{})
	s.addConversation(discrete.Carson{})
	s.addConversation(discrete.Carta{})
	s.addConversation(discrete.Casey{})
	s.addConversation(discrete.CenterOfTheMagicPentagram{})
	s.addConversation(discrete.Cesar3{})
	s.addConversation(discrete.Chef{})
	s.addConversation(discrete.Cherry{})
	s.addConversation(discrete.ChiefTatamo{})
	s.addConversation(discrete.Chris{})
	s.addConversation(discrete.ChunJi{})
	s.addConversation(discrete.Cliff{})
	s.addConversation(discrete.Cloy{})
	s.addConversation(discrete.Coco{})
	s.addConversation(discrete.ControlDevice{})
	s.addConversation(discrete.Corba{})
	s.addConversation(discrete.Corine{})
	s.addConversation(discrete.CornerOfTheMagicLibrary{})
	s.addConversation(discrete.CornerOfTheMagicLibrary2{})
	s.addConversation(discrete.Crane{})
	s.addConversation(discrete.CrawlsWithBalrog{})
	s.addConversation(discrete.CrumblingStatue{})
	s.addConversation(discrete.Cygnus{})
	s.addConversation(discrete.DancesWithBalrog{})
	s.addConversation(discrete.DancesWithBalrogDemo{})
	s.addConversation(discrete.DangerZoneTaxi{})
	s.addConversation(discrete.DarkLordDemo{})
	s.addConversation(discrete.DaveAndIris{})
	s.addConversation(discrete.DemonsDoorwayEastHenesys{})
	s.addConversation(discrete.DemonsDoorwayEllinia{})
	s.addConversation(discrete.DemonsDoorwayHenesys{})
	s.addConversation(discrete.DemonsDoorwayKerningCity{})
	s.addConversation(discrete.DemonsDoorwayPerion{})
	s.addConversation(discrete.Desk{})
	s.addConversation(discrete.DimensionalMirror{})
	s.addConversation(discrete.DocumentRoll{})
	s.addConversation(discrete.Dolphin{})
	s.addConversation(discrete.Dolphin2{})
	s.addConversation(discrete.DrFeeble{})
	s.addConversation(discrete.Dunamis{})
	s.addConversation(discrete.Duru{})
	s.addConversation(discrete.Eckhart{})
	s.addConversation(discrete.EckhartStatue{})
	s.addConversation(discrete.Egnet{})
	s.addConversation(discrete.Eleanor{})
	s.addConversation(discrete.Eleska{})
	s.addConversation(discrete.EllinForestMilepost{})
	s.addConversation(discrete.EncryptedSlateOfTheSquad{})
	s.addConversation(discrete.EnigmaTombstone{})
	s.addConversation(discrete.EntranceOfSealedShrine{})
	s.addConversation(discrete.Erin{})
	s.addConversation(discrete.FallenKnight{})
	s.addConversation(discrete.FirstEOSRock{})
	s.addConversation(discrete.FirstMagicPentagram{})
	s.addConversation(discrete.FirstPipeHandle{})
	s.addConversation(discrete.FountainOfLife{})
	s.addConversation(discrete.FourthEOSRock{})
	s.addConversation(discrete.Francis{})
	s.addConversation(discrete.Francois{})
	s.addConversation(discrete.Francis2{})
	s.addConversation(discrete.FranzTheOwner{})
	s.addConversation(discrete.Gate{})
	s.addConversation(discrete.Gaga{})
	s.addConversation(discrete.Geras{})
	s.addConversation(discrete.GhostHunterBob{})
	s.addConversation(discrete.Gina{})
	s.addConversation(discrete.GrandpaMoonBunny{})
	s.addConversation(discrete.GrendelTheReallyOld{})
	s.addConversation(discrete.GrendelTheReallyOldDemo{})
	s.addConversation(discrete.Gritto{})
	s.addConversation(discrete.HanTheBroker{})
	s.addConversation(discrete.Harmonia{})
	s.addConversation(discrete.HarpStringA{})
	s.addConversation(discrete.HarpStringB{})
	s.addConversation(discrete.HarpStringC{})
	s.addConversation(discrete.HarpStringD{})
	s.addConversation(discrete.HarpStringE{})
	s.addConversation(discrete.HarpStringF{})
	s.addConversation(discrete.HarpStringG{})
	s.addConversation(discrete.Harry{})
	s.addConversation(discrete.HarryCoconut{})
	s.addConversation(discrete.Hawkeye{})
	s.addConversation(discrete.HawkeyeStatue{})
	s.addConversation(discrete.Heena{})
	s.addConversation(discrete.Hellin{})
	s.addConversation(discrete.HenesysForest{})
	s.addConversation(discrete.HenesysForest2{})
	s.addConversation(discrete.Heracle{})
	s.addConversation(discrete.HiddenNote{})
	s.addConversation(discrete.HiddenNote2{})
	s.addConversation(discrete.Hikari{})
	s.addConversation(discrete.HolyStone{})
	s.addConversation(discrete.HotelReceptionist{})
	s.addConversation(discrete.HumanoidA{})
	s.addConversation(discrete.IcebyrdSlimm{})
	s.addConversation(discrete.IncompleteMagicSquare{})
	s.addConversation(discrete.InsignificantBeing{})
	s.addConversation(discrete.Irena{})
	s.addConversation(discrete.IrenaStatue{})
	s.addConversation(discrete.Irene{})
	s.addConversation(discrete.IsaTheStationGuide{})
	s.addConversation(discrete.Jack{})
	s.addConversation(discrete.Jake{})
	s.addConversation(discrete.Jane{})
	s.addConversation(discrete.Jano{})
	s.addConversation(discrete.Jean{})
	s.addConversation(discrete.Jeff{})
	s.addConversation(discrete.Jiyur{})
	s.addConversation(discrete.JMFromThaStreetz{})
	s.addConversation(discrete.Joel{})
	s.addConversation(discrete.JohnBarricade{})
	s.addConversation(discrete.Kanderune{})
	s.addConversation(discrete.Karcasa{})
	s.addConversation(discrete.Karen{})
	s.addConversation(discrete.Keeny{})
	s.addConversation(discrete.Kenta{})
	s.addConversation(discrete.Kerny{})
	s.addConversation(discrete.KingPepe{})
	s.addConversation(discrete.KinoKonoko{})
	s.addConversation(discrete.Kiridu{})
	s.addConversation(discrete.Kiriko{})
	s.addConversation(discrete.KiriruEreve{})
	s.addConversation(discrete.KiruEreve{})
	s.addConversation(discrete.KiruToOrbis{})
	s.addConversation(discrete.KiruOrbis{})
	s.addConversation(discrete.Kiruru{})
	s.addConversation(discrete.KnightArmor{})
	s.addConversation(discrete.Konpei{})
	s.addConversation(discrete.KonpeiExit{})
	s.addConversation(discrete.KonpeiExpeditionSuccessExit{})
	s.addConversation(discrete.KiriruEllinia{})
	s.addConversation(discrete.KyrinDemo{})
	s.addConversation(discrete.LePetitPrince{})
	s.addConversation(discrete.Lea{})
	s.addConversation(discrete.Legor{})
	s.addConversation(discrete.Lenario{})
	s.addConversation(discrete.Lila{})
	s.addConversation(discrete.Lira{})
	s.addConversation(discrete.Lisa{})
	s.addConversation(discrete.LitaLawless{})
	s.addConversation(discrete.LittleSuzy{})
	s.addConversation(discrete.Loha{})
	s.addConversation(discrete.Lohd{})
	s.addConversation(discrete.LordJonathan{})
	s.addConversation(discrete.Louis{})
	s.addConversation(discrete.Lukan{})
	s.addConversation(discrete.Luke{})
	s.addConversation(discrete.MachineApparatus{})
	s.addConversation(discrete.Maed{})
	s.addConversation(discrete.MagicianJobInstructor{})
	s.addConversation(discrete.MagicianJobInstructorExit{})
	s.addConversation(discrete.MagicianStatue{})
	s.addConversation(discrete.Maid{})
	s.addConversation(discrete.MapleLeafMarble{})
	s.addConversation(discrete.MarTheFairy{})
	s.addConversation(discrete.MarkTheToySoldier{})
	s.addConversation(discrete.Mel{})
	s.addConversation(discrete.MetalBucketSnowman{})
	s.addConversation(discrete.Meteorite1{})
	s.addConversation(discrete.Meteorite2{})
	s.addConversation(discrete.Meteorite3{})
	s.addConversation(discrete.Meteorite4{})
	s.addConversation(discrete.Meteorite5{})
	s.addConversation(discrete.Meteorite6{})
	s.addConversation(discrete.Mia{})
	s.addConversation(discrete.Mia2{})
	s.addConversation(discrete.Mihile{})
	s.addConversation(discrete.MihileStatue{})
	s.addConversation(discrete.Milla{})
	s.addConversation(discrete.Mimo{})
	s.addConversation(discrete.MinoTheOwner{})
	s.addConversation(discrete.Mo{})
	s.addConversation(discrete.MongFromKong{})
	s.addConversation(discrete.MonstrousLookingStatue{})
	s.addConversation(discrete.MoonstoneGrave{})
	s.addConversation(discrete.Moose{})
	s.addConversation(discrete.MooseExit{})
	s.addConversation(discrete.MotherMilkCow1{})
	s.addConversation(discrete.MotherMilkCow2{})
	s.addConversation(discrete.MrGoldstein{})
	s.addConversation(discrete.MrPickall{})
	s.addConversation(discrete.MrSmith{})
	s.addConversation(discrete.MrThunder{})
	s.addConversation(discrete.MsTan{})
	s.addConversation(discrete.MuLungDojoBulletinBoard{})
	s.addConversation(discrete.Mue{})
	s.addConversation(discrete.Muirhat{})
	s.addConversation(discrete.MysteriousStatue{})
	s.addConversation(discrete.Nara{})
	s.addConversation(discrete.Naran{})
	s.addConversation(discrete.Natalie{})
	s.addConversation(discrete.NautilusMidSizedTaxi{})
	s.addConversation(discrete.Neinheart{})
	s.addConversation(discrete.NeinheartDemo{})
	s.addConversation(discrete.Nella{})
	s.addConversation(discrete.Neru{})
	s.addConversation(discrete.NLCTaxi{})
	s.addConversation(discrete.Nuris{})
	s.addConversation(discrete.OldManTom{})
	s.addConversation(discrete.OlsonTheToySoldier{})
	s.addConversation(discrete.OrbisMagicSpot1{})
	s.addConversation(discrete.OrbisMagicSpot20{})
	s.addConversation(discrete.Oz{})
	s.addConversation(discrete.OzStatue{})
	s.addConversation(discrete.PalaceOasis{})
	s.addConversation(discrete.Papulatus{})
	s.addConversation(discrete.Parwen{})
	s.addConversation(discrete.Pason{})
	s.addConversation(discrete.Paul{})
	s.addConversation(discrete.Pelace{})
	s.addConversation(discrete.Perzen{})
	s.addConversation(discrete.Peter{})
	s.addConversation(discrete.Phil{})
	s.addConversation(discrete.Phyllia{})
	s.addConversation(discrete.PictureFrame{})
	s.addConversation(discrete.Pietra{})
	s.addConversation(discrete.Pietro{})
	s.addConversation(discrete.PirateJobInstructorExit{})
	s.addConversation(discrete.PirateStatue{})
	s.addConversation(discrete.Pison{})
	s.addConversation(discrete.PowerBForeBowman{})
	s.addConversation(discrete.PowerBForeEllinia{})
	s.addConversation(discrete.PowerBForeHenesys{})
	s.addConversation(discrete.PowerBForeKerningCity{})
	s.addConversation(discrete.PowerBForeMagician{})
	s.addConversation(discrete.PowerBForeNautilus{})
	s.addConversation(discrete.PowerBForePerion{})
	s.addConversation(discrete.PowerBForePirate{})
	s.addConversation(discrete.PowerBForeThief{})
	s.addConversation(discrete.PowerBForeWarrior{})
	s.addConversation(discrete.PracticeChart{})
	s.addConversation(discrete.PrinceGiuseppe{})
	s.addConversation(discrete.ProfessorFoxwit{})
	s.addConversation(discrete.Purin{})
	s.addConversation(discrete.PuroLithHarbor{})
	s.addConversation(discrete.PuroRien{})
	s.addConversation(discrete.PuroToLithHarbor{})
	s.addConversation(discrete.PuroToRien{})
	s.addConversation(discrete.QueensCabinet{})
	s.addConversation(discrete.Rain{})
	s.addConversation(discrete.Ramini{})
	s.addConversation(discrete.Refugee1{})
	s.addConversation(discrete.Refugee2{})
	s.addConversation(discrete.Refugee3{})
	s.addConversation(discrete.Refugee4{})
	s.addConversation(discrete.RegularCabEllinia{})
	s.addConversation(discrete.RegularCabHenesys{})
	s.addConversation(discrete.RegularCabKerningCity{})
	s.addConversation(discrete.RegularCabLithHarbor{})
	s.addConversation(discrete.RegularCabPerion{})
	s.addConversation(discrete.ReturningRock{})
	s.addConversation(discrete.Ria{})
	s.addConversation(discrete.Ridley{})
	s.addConversation(discrete.Rini{})
	s.addConversation(discrete.Rius{})
	s.addConversation(discrete.RobertHolly{})
	s.addConversation(discrete.Robin{})
	s.addConversation(discrete.RobinTheHuntress{})
	s.addConversation(discrete.RollyExit{})
	s.addConversation(discrete.Romi{})
	s.addConversation(discrete.Roodolph{})
	s.addConversation(discrete.Rooney{})
	s.addConversation(discrete.Rosey{})
	s.addConversation(discrete.Rupi{})
	s.addConversation(discrete.Russellon{})
	s.addConversation(discrete.RussellonsDesk{})
	s.addConversation(discrete.Sage{})
	s.addConversation(discrete.Sage2{})
	s.addConversation(discrete.Samuel{})
	s.addConversation(discrete.Santa2{})
	s.addConversation(discrete.ScarfSnowman{})
	s.addConversation(discrete.Schegerazade{})
	s.addConversation(discrete.SeargantAnderson{})
	s.addConversation(discrete.SecondEOSRock{})
	s.addConversation(discrete.SecondMagicPentagram{})
	s.addConversation(discrete.SecondPipeHandle{})
	s.addConversation(discrete.SecretWall{})
	s.addConversation(discrete.Sejan{})
	s.addConversation(discrete.Sera{})
	s.addConversation(discrete.Shadrion{})
	s.addConversation(discrete.Shalon{})
	s.addConversation(discrete.ShamanRock1{})
	s.addConversation(discrete.ShamanRock2{})
	s.addConversation(discrete.Shane{})
	s.addConversation(discrete.Shanks{})
	s.addConversation(discrete.SharenIIIsWill{})
	s.addConversation(discrete.Shawn{})
	s.addConversation(discrete.Shinsoo{})
	s.addConversation(discrete.ShinyStone{})
	s.addConversation(discrete.Shulynch{})
	s.addConversation(discrete.Shulynch2{})
	s.addConversation(discrete.SunstoneGrave{})
	s.addConversation(discrete.Shuri{})
	s.addConversation(discrete.Simon{})
	s.addConversation(discrete.Sion{})
	s.addConversation(discrete.Sirin{})
	s.addConversation(discrete.Slyn{})
	s.addConversation(discrete.SmallStreetLight{})
	s.addConversation(discrete.SmallTreeStump{})
	s.addConversation(discrete.SnowmanExit{})
	s.addConversation(discrete.SparklingCrystal{})
	s.addConversation(discrete.Spindle{})
	s.addConversation(discrete.Spinel{})
	s.addConversation(discrete.Spiruna{})
	s.addConversation(discrete.StaffSeargantCharlie{})
	s.addConversation(discrete.Steward{})
	s.addConversation(discrete.StrangeLookingStatue{})
	s.addConversation(discrete.StrawHatSnowman{})
	s.addConversation(discrete.SubwayExit{})
	s.addConversation(discrete.SubwayTrashCan1{})
	s.addConversation(discrete.SubwayTrashCan2{})
	s.addConversation(discrete.SubwayTrashCan3{})
	s.addConversation(discrete.SubwayTrashCan4{})
	s.addConversation(discrete.Sunny{})
	s.addConversation(discrete.Syras{})
	s.addConversation(discrete.T1337{})
	s.addConversation(discrete.Taggrin{})
	s.addConversation(discrete.Tangyoon{})
	s.addConversation(discrete.Tess{})
	s.addConversation(discrete.TheGlimmerManNLC{})
	s.addConversation(discrete.TheTicketGate{})
	s.addConversation(discrete.ThiefJobInstructor{})
	s.addConversation(discrete.ThiefJobInstructorExit{})
	s.addConversation(discrete.ThiefStatue{})
	s.addConversation(discrete.ThirdEOSRock{})
	s.addConversation(discrete.ThirdMagicPentagram{})
	s.addConversation(discrete.ThirdPipeHandle{})
	s.addConversation(discrete.ThomasSwift{})
	s.addConversation(discrete.ThreeRefugees{})
	s.addConversation(discrete.Tia{})
	s.addConversation(discrete.Tian{})
	s.addConversation(discrete.Tigun{})
	s.addConversation(discrete.TimeGate{})
	s.addConversation(discrete.Tombstone{})
	s.addConversation(discrete.Tommie{})
	s.addConversation(discrete.Tony{})
	s.addConversation(discrete.TrainerBartos{})
	s.addConversation(discrete.TrainerFrod{})
	s.addConversation(discrete.TrashCan{})
	s.addConversation(discrete.TreasureChestB1{})
	s.addConversation(discrete.TreasureChestB2{})
	s.addConversation(discrete.TreasureChestB3{})
	s.addConversation(discrete.TutorialLilin{})
	s.addConversation(discrete.TylusComplete{})
	s.addConversation(discrete.Vicious{})
	s.addConversation(discrete.Victoria{})
	s.addConversation(discrete.Vikin{})
	s.addConversation(discrete.VIPCabEllinia{})
	s.addConversation(discrete.VIPCabLithHarbor{})
	s.addConversation(discrete.WarriorJobInstructor{})
	s.addConversation(discrete.WarriorJobInstructorExit{})
	s.addConversation(discrete.WarriorStatue{})
	s.addConversation(discrete.WaterFilter{})
	s.addConversation(discrete.Weaver{})
	s.addConversation(discrete.Wisp{})
	s.addConversation(discrete.WitchMalady{})
	s.addConversation(discrete.WolfGuard{})
	s.addConversation(discrete.WolfSpiritRyko{})
	s.addConversation(discrete.Xinga{})
	s.addConversation(discrete.YokoYoko{})
	s.addConversation(discrete.YoungAthenaPierce{})
	s.addConversation(discrete.YoungAthenaPierce2{})
	s.addConversation(discrete.YuleteDefeated{})
	s.addConversation(discrete.Yuris{})
	return s
}

func (s *Registry) GetScript(npcId uint32) (*script.Script, error) {
	if val, ok := s.registry[npcId]; ok {
		return &val, nil
	}
	return nil, errors.New("unable to locate script")
}

func (s *Registry) addConversation(handler script.Script) {
	s.registry[handler.NPCId()] = handler
}
