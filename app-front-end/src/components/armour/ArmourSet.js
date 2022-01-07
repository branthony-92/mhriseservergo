import { useState } from "react";
import Card from "../ui/Card";
import ArmourPiece from "./ArmourPiece";
import ArmourSetSummary from "./ArmourSetSummary";
import ArmourStats from "./ArmourStats";

const helmImgURL = "https://fanonmonsterhunter.fandom.com/wiki/Armor_Icons?file=Helmet_Icon_White.png";
const mailImgURL = "https://fanonmonsterhunter.fandom.com/wiki/Armor_Icons?file=Chest_Icon_White.png";
const vambracesImgURL = "https://fanonmonsterhunter.fandom.com/wiki/Armor_Icons?file=Arm_Icon_White.png";
const coilImgURL = "https://fanonmonsterhunter.fandom.com/wiki/Armor_Icons?file=Waist_Icon_White.png";
const gravesImgURL = "https://fanonmonsterhunter.fandom.com/wiki/Armor_Icons?file=Leg_Icon_White.png";

function ArmourSet(props) {
    const [ showStates, setShowStates ] = useState(false);

    function handleToggleStats() {
        console.log("clicked")
        setShowStates(!showStates)
    }

    function generateStateSummary(set) {
        let stats = {
            defence: 0,              
            fire_res: 0,             
            water_res: 0,           
            thunder_res: 0,           
            ice_res: 0,    
            dragon_res: 0 
        };

        if (set.helm) {
            stats.defence += set.helm.defence;              
            stats.fire_res += set.helm.fire_res;  
            stats.water_res += set.helm.water_res;  
            stats.thunder_res += set.helm.thunder_res;  
            stats.ice_res += set.helm.ice_res;  
            stats.dragon_res += set.helm.dragon_res;  
        }
        if (set.mail) {
            stats.defence += set.mail.defence;              
            stats.fire_res += set.mail.fire_res;  
            stats.water_res += set.mail.water_res;  
            stats.thunder_res += set.mail.thunder_res;  
            stats.ice_res += set.mail.ice_res;  
            stats.dragon_res += set.mail.dragon_res;  
        }
        if (set.coil) {
            stats.defence += set.coil.defence;              
            stats.fire_res += set.coil.fire_res;  
            stats.water_res += set.coil.water_res;  
            stats.thunder_res += set.coil.thunder_res;  
            stats.ice_res += set.coil.ice_res;  
            stats.dragon_res += set.coil.dragon_res;  
        }
        if (set.vambraces) {
            stats.defence += set.vambraces.defence;              
            stats.fire_res += set.vambraces.fire_res;  
            stats.water_res += set.vambraces.water_res;  
            stats.thunder_res += set.vambraces.thunder_res;  
            stats.ice_res += set.vambraces.ice_res;  
            stats.dragon_res += set.vambraces.dragon_res;  
        }
        if (set.greaves) {
            stats.defence += set.greaves.defence;              
            stats.fire_res += set.greaves.fire_res;  
            stats.water_res += set.greaves.water_res;  
            stats.thunder_res += set.greaves.thunder_res;  
            stats.ice_res += set.greaves.ice_res;  
            stats.dragon_res += set.greaves.dragon_res;  
        }
        return stats;
    }

    return (
        <li key={props.key}> 
            <Card>
                <h2 onClick={handleToggleStats}>{props.setName}</h2>
                {showStates && <ArmourSetSummary key={props.key} info={generateStateSummary(props)}/>}
                {showStates && <ArmourPiece key={props.key} img={helmImgURL} info={props.helm}/>}
                {showStates && <ArmourPiece key={props.key} img={mailImgURL} info={props.mail} />}
                {showStates && <ArmourPiece key={props.key} img={vambracesImgURL} info={props.vambraces} />}
                {showStates && <ArmourPiece key={props.key} img={coilImgURL} info={props.coil} />}
                {showStates && <ArmourPiece key={props.key} img={gravesImgURL} info={props.greaves} />}
            </Card>
        </li>
    );
}

export default ArmourSet;