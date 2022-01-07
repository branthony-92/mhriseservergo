import { useState } from "react";
import Card from "../ui/Card";
import ArmourStats from "./ArmourStats";

function ArmourPiece(props) {

    const [ showStates, setShowStates ] = useState(false);

    function handleToggleStats() {
        console.log("clicked")
        setShowStates(!showStates)
    }

    if (props.info != null) {
        return (
            <Card>
                <div onClick={handleToggleStats}>
                <h3>{props.info.piece_name}</h3>   
                    { showStates && <ArmourStats info={props.info}/>}
                    { showStates && <image src={props.img} alt={props.info.piece_name}/> }
                </div>
            </Card>
        );
    } else {
        return <div/>
    }

}

export default ArmourPiece;