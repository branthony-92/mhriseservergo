import { useState, useEffect } from "react";

import ArmourList from "../components/armour/ArmourList";

const DEMO_SET = [
    {
        helm: {
            _id: 0,
            piece_name: "Aelucanth Vertex S",
            set_name: "Aelucanth S Set",
            piece_type: "helm",
            rarity: 6,
            defence: 62,
            fire_res: -2,
            water_res: 0,
            thunder_res: 3,
            ice_res: -2,
            dragon_res: 2,
            skills: [
                {
                    name: "Critical Element",
                    level: 2
                }
            ]
        },
        mail: {
            _id: 1,
            piece_name: "Aelucanth Thorax S",
           set_name: "Aelucanth S Set",
            piece_type: "mail",
            rarity: 6,
            defence: 62,
            fire_res: -2,
            water_res: 0,
            thunder_res: 3,
            ice_res: -2,
            dragon_res: 2,
            skills: [
                {
                    name: "Critical Element",
                    level: 1
                },
                {
                    name: "Critical Eye",
                    level: 2
                },
                {
                    name: "Dragon Attack",
                    level: 1
                }
            ]
        },
        coil: {
            _id: 2,
            piece_name: "Aelucanth Elytra S",
            set_name: "Aelucanth S Set",
            piece_type: "coil",
            rarity: 6,
            defence: 62,
            fire_res: -2,
            water_res: 0,
            thunder_res: 3,
            ice_res: -2,
            dragon_res: 2,
            skills: [
                {
                    name: "Dragon Attack",
                    level: 3
                }
            ]
        },
        vambraces: {
            _id: 3,
            piece_name: "Aelucanth Brachia S",
            set_name: "Aelucanth S Set",
            piece_type: "vambraces",
            rarity: 6,
            defence: 62,
            fire_res: -2,
            water_res: 0,
            thunder_res: 3,
            ice_res: -2,
            dragon_res: 2,
            skills: [
                {
                    name: "Evade Window",
                    level: 1
                }
            ]
        },
        greaves: {
            _id: 4,
            piece_name: "Aelucanth Crura S",
            set_name: "Aelucanth S Set",
            piece_type: "greaves",
            rarity: 6,
            defence: 62,
            fire_res: -2,
            water_res: 0,
            thunder_res: 3,
            ice_res: -2,
            dragon_res: 2,
            skills: [
                {
                    name: "Critical Eye",
                    level: 2
                },
                {
                    name: "Dragon Attack",
                    level: 1
                }
            ]
        },
        set_name: "Aelucanth S Set",
        summary: ""
    }
    ,
    {
    helm: {
        _id: 5,
        piece_name: "Bullfango Mask S",
        set_name: "Bullfango S Set",
        piece_type: "helm",
        rarity: 4,
        defence: 36,
        fire_res: 0,
        water_res: 4,
        thunder_res: -2,
        ice_res: 0,
        dragon_res: 0,
        skills: [
            {
                name: "Bludgeoner",
                level: 2
            }
        ]
    },
    mail: null,
    coil: null,
    vambraces: null,
    greaves: null,
    set_name: "Bullfango S Set",
    summary: ""
}
];

function ArmourSetsPage(props) {
    const [isLoading, setIsLoading] = useState(true);
    const [loadedArmour, setLoadedArmour] = useState([]);

    // load the armour for the page
    useEffect(() => {
        setIsLoading(true);
        console.log("fetching");
        fetch(
            'http://localhost:4000/api/v1/armoursets'
        ).then(response => {
            console.log("waiting for response")
            console.log(response)
            return response.json();
        }).then(data => {
            console.log("reading data")
            setIsLoading(false)
            setLoadedArmour(data.message_body)
        });
    }, []);

    if (isLoading) {
        return (
            <section>
                <p>Loading...</p>
            </section>
        );
    }
    return (
        <section>
            <ArmourList armourSets={loadedArmour}/>
        </section>
    );
}

export default ArmourSetsPage;