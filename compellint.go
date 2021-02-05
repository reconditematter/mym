// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

import "math"

// completeBD -- computes auxiliary complete elliptic integrals B(m), D(m).
// Reference: Toshio Fukushima, Math. Comp. 80 (2011), 1725-1743,
// Precise and fast computation of the general complete elliptic integral of the second kind.
func completeBD(m float64) (B, D float64) {
	const (
		B1 = 1.0 / 2.0
		B2 = 1.0 / 16.0
		B3 = 3.0 / 128.0
		B4 = 25.0 / 2048.0
		B5 = 245.0 / 32768.0
		B6 = 1323.0 / 262144.0
		B7 = 7623.0 / 2097152.0
		B8 = 184041.0 / 67108864.0
	)
	const (
		D1 = 1.0 / 2.0
		D2 = 3.0 / 16.0
		D3 = 15.0 / 128.0
		D4 = 175.0 / 2048.0
		D5 = 2205.0 / 32768.0
		D6 = 14553.0 / 262144.0
		D7 = 99099.0 / 2097152.0
		D8 = 2760615.0 / 67108864.0
	)
	const (
		Q1  = 1.0 / 16.0
		Q2  = 1.0 / 32.0
		Q3  = 21.0 / 1024.0
		Q4  = 31.0 / 2048.0
		Q5  = 6257.0 / 524288.0
		Q6  = 10293.0 / 1048576.0
		Q7  = 279025.0 / 33554432.0
		Q8  = 483127.0 / 67108864.0
		Q9  = 435506703.0 / 68719476736.0
		Q10 = 776957575.0 / 137438953472.0
		Q11 = 22417045555.0 / 4398046511104.0
		Q12 = 40784671953.0 / 8796093022208.0
		Q13 = 9569130097211.0 / 2251799813685248.0
		Q14 = 17652604545791.0 / 4503599627370496.0
		Q15 = 523910972020563.0 / 144115188075855872.0
		Q16 = 976501268709949.0 / 288230376151711744.0
	)
	const (
		K1 = 1.0 / 4.0
		K2 = 9.0 / 64.0
		K3 = 25.0 / 256.0
		K4 = 1225.0 / 16384.0
		K5 = 3969.0 / 65536.0
		K6 = 53361.0 / 1048576.0
		K7 = 184041.0 / 4194304.0
	)
	mc := 1 - m
	switch {
	case m <= Epsilon:
		B = math.Pi / 4
		D = math.Pi / 4
	case m <= 0.01:
		B = (math.Pi / 2) * (B1 + m*(B2+m*(B3+m*(B4+m*(B5+m*(B6+m*(B7+m*B8)))))))
		D = (math.Pi / 2) * (D1 + m*(D2+m*(D3+m*(D4+m*(D5+m*(D6+m*(D7+m*D8)))))))
	case m <= 0.1:
		mx := 0.95 - mc
		B = 0.790401413584395132310045630540381158921005 + mx*(0.102006266220019154892513446364386528537788+mx*(0.039878395558551460860377468871167215878458+mx*(0.021737136375982167333478696987134316809322+mx*(0.013960979767622057852185340153691548520857+mx*(0.009892518822669142478846083436285145400444+mx*(0.007484612400663335676130416571517444936951+mx*(0.005934625664295473695080715589652011420808+mx*(0.004874249053581664096949448689997843978535+mx*(0.004114606930310886136960940893002069423559+mx*(0.003550452989196176932747744728766021440856+mx*(0.003119229959988474753291950759202798352266)))))))))))
		D = 0.800602040206397047799296975176819811774784 + mx*(0.313994477771767756849615832867393028789057+mx*(0.205913118705551954501930953451976374435088+mx*(0.157744346538923994475225014971416837073598+mx*(0.130595077319933091909091103101366509387938+mx*(0.113308474489758568672985167742047066367053+mx*(0.101454199173630195376251916342483192174927+mx*(0.0929187842072974367037702927967784464949434+mx*(0.0865653801481680871714054745336652101162894+mx*(0.0817279846651030135350056216958053404884715+mx*(0.0779906657291070378163237851392095284454654+mx*(0.075080426851268007156477347905308063808848)))))))))))
	case m <= 0.2:
		mx := 0.85 - mc
		B = 0.80102406445284489393880821604009991524037 + mx*(0.11069534452963401497502459778015097487115+mx*(0.047348746716993717753569559936346358937777+mx*(0.028484367255041422845322166419447281776162+mx*(0.020277811444003597057721308432225505126013+mx*(0.015965005853099119442287313909177068173564+mx*(0.013441320273553634762716845175446390822633+mx*(0.011871565736951439501853534319081030547931+mx*(0.010868363672485520630005005782151743785248+mx*(0.010231587232710564565903812652581252337697+mx*(0.009849585546666211201566987057592610884309+mx*(0.009656606347153765129943681090056980586986)))))))))))
		D = 0.834232667811735098431315595374145207701720 + mx*(0.360495281619098275577215529302260739976126+mx*(0.262379664114505869328637749459234348287432+mx*(0.223723944518094276386520735054801578584350+mx*(0.206447811775681052682922746753795148394463+mx*(0.199809440876486856438050774316751253389944+mx*(0.199667451603795274869211409350873244844882+mx*(0.204157558868236842039815028663379643303565+mx*(0.212387467960572375038025392458549025660994+mx*(0.223948914061499360356873401571821627069173+mx*(0.238708097425597860161720875806632864507536+mx*(0.256707203545463755643710021815937785120030)))))))))))
	case m <= 0.3:
		mx := 0.75 - mc
		B = 0.81259777291992049322557009456643357559904 + mx*(0.12110961794551011284012693733241967660542+mx*(0.057293376831239877456538980381277010644332+mx*(0.038509451602167328057004166642521093142114+mx*(0.030783430301775232744816612250699163538318+mx*(0.027290564934732526869467118496664914274956+mx*(0.025916369289445198731886546557337255438083+mx*(0.025847203343361799141092472018796130324244+mx*(0.026740923539348854616932735567182946385269+mx*(0.028464314554825704963640157657034405579849+mx*(0.030995446237278954096189768338119395563447+mx*(0.034384369179940975864103666824736551261799+mx*(0.038738002072493935952384233588242422046537))))))))))))
		D = 0.873152581892675549645633563232643413901757 + mx*(0.420622230667770215976919792378536040460605+mx*(0.344231061559450379368201151870166692934830+mx*(0.331133021818721761888662390999045979071436+mx*(0.345277285052808411877017306497954757532251+mx*(0.377945322150393391759797943135325823338761+mx*(0.427378012464553880508348757311170776829930+mx*(0.494671744307822405584118022550673740404732+mx*(0.582685115665646200824237214098764913658889+mx*(0.695799207728083164790111837174250683834359+mx*(0.840018401472533403272555302636558338772258+mx*(1.023268503573606060588689738498395211300483+mx*(1.255859085136282496149035687741403985044122))))))))))))
	case m <= 0.4:
		mx := 0.65 - mc
		B = 0.8253235579835158949845697805395190063745 + mx*(0.1338621160836877898575391383950840569989+mx*(0.0710112935979886745743770664203746758134+mx*(0.0541784774173873762208472152701393154906+mx*(0.0494517449481029932714386586401273353617+mx*(0.0502221962241074764652127892365024315554+mx*(0.0547429131718303528104722303305931350375+mx*(0.0627462579270016992000788492778894700075+mx*(0.0746698810434768864678760362745179321956+mx*(0.0914808451777334717996463421986810092918+mx*(0.1147050921109978235104185800057554574708+mx*(0.1465711325814398757043492181099197917984+mx*(0.1902571373338462844225085057953823854177))))))))))))
		D = 0.9190270392420973478848471774160778462738 + mx*(0.5010021592882475139767453081737767171354+mx*(0.4688312705664568629356644841691659415972+mx*(0.5177142277764000147059587510833317474467+mx*(0.6208433913173031070711926900889045286988+mx*(0.7823643937868697229213240489900179142670+mx*(1.0191145350761029126165253557593691585239+mx*(1.3593452027484960522212885423056424704073+mx*(1.8457173023588279422916645725184952058635+mx*(2.5410717031539207287662105618152273788399+mx*(3.5374046552080413366422791595672470037341+mx*(4.9692960029774259303491034652093672488707+mx*(7.0338228700300311264031522795337352226926+mx*(10.020043225034471401553194050933390974016)))))))))))))
	case m <= 0.5:
		mx := 0.55 - mc
		B = 0.8394795702706129706783934654948360410325 + mx*(0.1499164403063963359478614453083470750543+mx*(0.0908319358194288345999005586556105610069+mx*(0.0803470334833417864262134081954987019902+mx*(0.0856384405004704542717663971835424473169+mx*(0.1019547259329903716766105911448528069506+mx*(0.1305748115336160150072309911623351523284+mx*(0.1761050763588499277679704537732929242811+mx*(0.2468351644029554468698889593583314853486+mx*(0.3564244768677188553323196975301769697977+mx*(0.5270025622301027434418321205779314762241+mx*(0.7943896342593047502260866957039427731776+mx*(1.2167625324297180206378753787253096783993))))))))))))
		D = 0.9744043665463696730314687662723484085813 + mx*(0.6132468053941609101234053415051402349752+mx*(0.6710966695021669963502789954058993004082+mx*(0.8707276201850861403618528872292437242726+mx*(1.2295422312026907609906452348037196571302+mx*(1.8266059675444205694817638548699906990301+mx*(2.8069345309977627400322167438821024032409+mx*(4.4187893290840281339827573139793805587268+mx*(7.0832360574787653249799018590860687062869+mx*(11.515088120557582942290563338274745712174+mx*(18.931511185999274639516732819605594455165+mx*(31.411996938204963878089048091424028309798+mx*(52.520729454575828537934780076768577185134+mx*(88.384854735065298062125622417251073520996+mx*(149.56637449398047835236703116483092644714+mx*(254.31790843104117434615624121937495622372)))))))))))))))
	case m <= 0.6:
		mx := 0.45 - mc
		B = 0.8554696151564199914087224774321783838373 + mx*(0.1708960726897395844132234165994754905373+mx*(0.1213352290269482260207667564010437464156+mx*(0.1282018835749474096272901529341076494573+mx*(0.1646872814515275597348427294090563472179+mx*(0.2374189087493817423375114793658754489958+mx*(0.3692081047164954516884561039890508294508+mx*(0.6056587338479277173311618664015401963868+mx*(1.0337055615578127436826717513776452476106+mx*(1.8189884893632678849599091011718520567105+mx*(3.2793776512738509375806561547016925831128+mx*(6.0298883807175363312261449542978750456611+mx*(11.269796855577941715109155203721740735793+mx*(21.354577850382834496786315532111529462693)))))))))))))
		D = 1.04345529511513353426326823569160142342838 + mx*(0.77962572192850485048535711388072271372632+mx*(1.02974236093206758187389128668777397528702+mx*(1.62203722341135313022433907993860147395972+mx*(2.78798953118534762046989770119382209443756+mx*(5.04838148737206914685643655935236541332892+mx*(9.46327761194348429539987572314952029503864+mx*(18.1814899494276679043749394081463811247757+mx*(35.5809805911791687037085198750213045708148+mx*(70.6339354619144501276254906239838074917358+mx*(141.828580083433059297030133195739832297859+mx*(287.448751250132166257642182637978103762677+mx*(587.115384649923076181773192202238389711345+mx*(1207.06543522548061603657141890778290399603+mx*(2495.58872724866422273012188618178997342537+mx*(5184.69242939480644062471334944523925163600+mx*(10817.2133369041327524988910635205356016939))))))))))))))))
	case m <= 0.7:
		mx := 0.35 - mc
		B = 0.8739200618486431359820482173294324246058 + mx*(0.1998140574823769459497418213885348159654+mx*(0.1727696158780152128147094051876565603862+mx*(0.2281069132842021671319791750725846795701+mx*(0.3704681411180712197627619157146806221767+mx*(0.6792712528848205545443855883980014994450+mx*(1.3480084966817573020596179874311042267679+mx*(2.8276709768538207038046918250872679902352+mx*(6.1794682501239140840906583219887062092430+mx*(13.935686010342811497608625663457407447757+mx*(32.218929281059722026322932181420383764028+mx*(76.006962959226101026399085304912635262362+mx*(182.32144908775406957609058046006949657416+mx*(443.51507644112648158679360783118806161062+mx*(1091.8547229028388292980623647414961662223+mx*(2715.7658664038195881056269799613407111521)))))))))))))))
		D = 1.13367833657573316571671258513452768536080 + mx*(1.04864317372997039116746991765351150490010+mx*(1.75346504119846451588826580872136305225406+mx*(3.52318272680338551269021618722443199230946+mx*(7.74947641381397458240336356601913534598302+mx*(17.9864500558507330560532617743406294626849+mx*(43.2559163462326133313977294448984936591235+mx*(106.681534454096017031613223924991564429656+mx*(268.098486573117433951562111736259672695883+mx*(683.624114850289804796762005964155730439745+mx*(1763.49708521918740723028849567007874329637+mx*(4592.37475383116380899419201719007475759114+mx*(12053.4410190488892782190764838488156555734+mx*(31846.6630207420816960681624497373078887317+mx*(84621.2213590568080177035346867495326879117+mx*(225956.423182907889987641304430180593010940+mx*(605941.517281758859958050194535269219533685+mx*(1.63108259953926832083633544697688841456604e6)))))))))))))))))
	case m <= 0.8:
		mx := 0.25 - mc
		B = 0.895902820924731621258525533131864225704 + mx*(0.243140003766786661947749288357729051637+mx*(0.273081875594105531575351304277604081620+mx*(0.486280007533573323895498576715458103274+mx*(1.082747437228230914750752674136983406683+mx*(2.743445290986452500459431536349945437824+mx*(7.555817828670234627048618342026400847824+mx*(22.05194082493752427472777448620986154515+mx*(67.15640644740229407624192175802742979626+mx*(211.2722537881770961691291434845898538537+mx*(681.9037843053270682273212958093073895805+mx*(2246.956231592536516768812462150619631201+mx*(7531.483865999711792004783423815426725079+mx*(25608.51260130241579018675054866136922157+mx*(88140.74740089604971425934283371277143256+mx*(306564.4242098446591430938434419151070722+mx*(1.076036077811072193752770590363885180738e6+mx*(3.807218502573632648224286313875985190526e6+mx*(1.356638224422139551020110323739879481042e7))))))))))))))))))
		D = 1.26061282657491161418014946566845780315983 + mx*(1.54866563808267658056930177790599939977154+mx*(3.55366941187160761540650011660758187283401+mx*(9.90044467610439875577300608183010716301714+mx*(30.3205666174524719862025105898574414438275+mx*(98.1802586588830891484913119780870074464833+mx*(329.771010434557055036273670551546757245808+mx*(1136.65598974289039303581967838947708073239+mx*(3993.83433574622979757935610692842933356144+mx*(14242.7295865552708506232731633468180669284+mx*(51394.7572916887209594591528374806790960057+mx*(187246.702914623152141768788258141932569037+mx*(687653.092375389902708761221294282367947659+mx*(2.54238553565398227033448846432182516906624e6+mx*(9.45378121934749027243313241962076028066811e6+mx*(3.53283630179709170835024033154326126569613e7+mx*(1.32593262383393014923560730485845833322771e8+mx*(4.99544968184054821463279808395426941549833e8+mx*(1.88840934729443872364972817525484292678543e9+mx*(7.16026753447893719179055010636502508063102e9+mx*(2.72233079469633962247554894093591262281929e10))))))))))))))))))))
	case m <= 0.85:
		mx := 0.175 - mc
		B = 0.915922052601931494319853880201442948834592 + mx*(0.294714252429483394379515488141632749820347+mx*(0.435776709264636140422971598963772380161131+mx*(1.067328246493644238508159085364429570207744+mx*(3.327844118563268085074646976514979307993733+mx*(11.90406004445092906188837729711173326621810+mx*(46.47838820224626393512400481776284680677096+mx*(192.7556002578809476962739389101964074608802+mx*(835.3356299261900063712302517586717381557137+mx*(3743.124548343029102644419963712353854902019+mx*(17219.07731004063094108708549153310467326395+mx*(80904.60401669850158353080543152212152282878+mx*(386808.3292751742460123683674607895217760313+mx*(1.876487670110449342170327796786290400635732e6+mx*(9.216559908641567755240142886998737950775910e6))))))))))))))
		D = 1.402200569110579095046054435635136986038164 + mx*(2.322205897861749446534352741005347103992773+mx*(7.462158366466719682730245467372788273333992+mx*(29.43506890797307903104978364254987042421285+mx*(128.1590924337895775262509354898066132182429+mx*(591.0807036911982326384997979640812493154316+mx*(2830.546229607726377048576057043685514661188+mx*(13917.76431889392229954434840686741305556862+mx*(69786.10525163921228258055074102587429394212+mx*(355234.1420341879634781808899208309503519936+mx*(1.830019186413931053503912913904321703777885e6+mx*(9.519610812032515607466102200648641326190483e6+mx*(4.992086875574849453986274042758566713803723e7+mx*(2.635677009826023473846461512029006874800883e8+mx*(1.399645765120061118824228996253541612110338e9+mx*(7.469935792837635004663183580452618726280406e9+mx*(4.004155595835610574316003488168804738481448e10+mx*(2.154630668144966654449602981243932210695662e11)))))))))))))))))
	case m <= 0.9:
		mx := 0.125 - mc
		B = 0.931906061029524827613331428871579482766771 + mx*(0.348448029538453860999386797137074571589376+mx*(0.666809178846938247558793864839434184202736+mx*(2.210769135708128662563678717558631573758222+mx*(9.491765048913406881414290930355300611703187+mx*(47.09304791027740853381457907791343619298913+mx*(255.9200460211233087050940506395442544885608+mx*(1480.029532675805407554800779436693505109703+mx*(8954.040904734313578374783155553041065984547+mx*(56052.48220982686949967604699243627330816542+mx*(360395.7241626000916973524840479780937869149+mx*(2.367539415273216077520928806581689330885103e6+mx*(1.582994957277684102454906900025484391190264e7+mx*(1.074158093278511100137056972128875270067228e8+mx*(7.380585460239595691878086073095523043390649e8+mx*(5.126022002555101496684687154904781856830296e9+mx*(3.593534065502416588712409180013118409428367e10+mx*(2.539881257612812212004146637239987308133582e11+mx*(1.808180007145359569674767150594344316702507e12))))))))))))))))))
		D = 1.541690112721819084362258323861459983048179 + mx*(3.379176214579645449453938918349243359477706+mx*(14.94058385670236671625328259137998668324435+mx*(81.91773929235074880784578753539752529822986+mx*(497.4900546551479866036061853049402721939835+mx*(3205.184010234846235275447901572262470252768+mx*(21457.32237355321925571253220641357074594515+mx*(147557.0156564174712105689758692929775004292+mx*(1.035045290185256525452269053775273002725343e6+mx*(7.371922334832212125197513363695905834126154e6+mx*(5.314344395142401141792228169170505958906345e7+mx*(3.868823475795976312985118115567305767603128e8+mx*(2.839458401528033778425531336599562337200510e9+mx*(2.098266122943898941547136470383199468548861e10+mx*(1.559617754017662417944194874282275405676282e11+mx*(1.165096220419884791236699872205721392201682e12+mx*(8.742012983013913804987431275193291316808766e12+mx*(6.584725462672366918676967847406180155459650e13+mx*(4.976798737062434393396993620379481464465749e14+mx*(3.773018634056605404718444239040628892506293e15+mx*(2.868263194837819660109735981973458220407767e16))))))))))))))))))))
	default:
		if mc <= Epsilon {
			B = 1.0
			D = (2*math.Ln2 - 1) - 0.5*math.Log(mc)
		} else {
			nome := mc * (Q1 + mc*(Q2+mc*(Q3+mc*(Q4+mc*(Q5+mc*(Q6+mc*(Q7+mc*(Q8+mc*(Q9+mc*(Q10+mc*(Q11+mc*(Q12+mc*(Q13+mc*(Q14+mc*(Q15+mc*Q16)))))))))))))))
			var dkkc, dddc float64
			if mc <= 0.01 {
				dkkc = mc * (K1 + mc*(K2+mc*(K3+mc*(K4+mc*(K5+mc*(K6+mc*K7))))))
				dddc = mc * (D1 + mc*(D2+mc*(D3+mc*(D4+mc*(D5+mc*(D6+mc*D7))))))
			} else {
				mx := mc - 0.05
				dkkc = 0.01286425658832983978282698630501405107893 + mx*(0.26483429894479586582278131697637750604652+mx*(0.15647573786069663900214275050014481397750+mx*(0.11426146079748350067910196981167739749361+mx*(0.09202724415743445309239690377424239940545+mx*(0.07843218831801764082998285878311322932444+mx*(0.06935260142642158347117402021639363379689+mx*(0.06293203529021269706312943517695310879457+mx*(0.05821227592779397036582491084172892108196+mx*(0.05464909112091564816652510649708377642504+mx*(0.05191068843704411873477650167894906357568+mx*(0.04978344771840508342564702588639140680363+mx*(0.04812375496807025605361215168677991360500))))))))))))
				dddc = 0.02548395442966088473597712420249483947953 + mx*(0.51967384324140471318255255900132590084179+mx*(0.20644951110163173131719312525729037023377+mx*(0.13610952125712137420240739057403788152260+mx*(0.10458014040566978574883406877392984277718+mx*(0.08674612915759188982465635633597382093113+mx*(0.07536380269617058326770965489534014190391+mx*(0.06754544594618781950496091910264174396541+mx*(0.06190939688096410201497509102047998554900+mx*(0.05771071515451786553160533778648705873199+mx*(0.05451217098672207169493767625617704078257+mx*(0.05204028407582600387265992107877094920787+mx*(0.05011532514520838441892567405879742720039))))))))))))
			}
			kkc := 1.0 + dkkc
			logq2 := -0.5 * math.Log(nome)
			elk := kkc * logq2
			dele := -dkkc/kkc + logq2*dddc
			elk1 := elk - 1.0
			delb := (dele - mc*elk1) / m
			B = 1.0 + delb
			D = elk1 - delb
		}
	}
	return
}

// Complete12 -- computes the complete elliptic integrals of the first
// and second kinds of the elliptic parameter `m`.
// This function returns (NaN,NaN) when m∉[0,1].
func Complete12(m float64) (K, E float64) {
	if !(0 <= m && m <= 1) {
		K, E = math.NaN(), math.NaN()
		return
	}
	if m <= Tiny {
		K, E = math.Pi/2, math.Pi/2
		return
	}
	if m == 1 {
		K, E = math.Inf(1), 1
		return
	}
	//
	B, D := completeBD(m)
	K = B + D
	E = B + (1-m)*D
	return
}
