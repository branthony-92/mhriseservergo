function ArmourSetSummary(props) {
    return (
        <div>
              <section>    
                <h3>Total Set Stats</h3>               
                <hr/>
                <p>Defense: <b>{props.info.defence}</b></p>               
                <p>Fire Res: <b>{props.info.fire_res}</b></p>               
                <p>Water Res: <b>{props.info.water_res}</b></p>            
                <p>Thunder Res: <b>{props.info.thunder_res}</b></p>            
                <p>Ice Res: <b>{props.info.ice_res}</b></p>             
                <p>Dragon Res: <b>{props.info.dragon_res}</b></p>       
            </section>
        </div>
    )
}

export default ArmourSetSummary