import ArmourSet from "./ArmourSet";

function ArmourList(props) {
    return (
        <ul>
            {props.armourSets.map((set =>
                <ArmourSet  
                key={set.key}    
                setName={set.set_name}       
                helm={set.helm}
                mail={set.mail}
                vambraces={set.vambraces}
                coil={set.coil}
                greaves={set.greaves}
                /> 
            ))}
        </ul>
    );
}

export default ArmourList;