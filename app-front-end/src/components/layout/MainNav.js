import {Link} from 'react-router-dom'
import classes from './MainNav.module.css'

function MainNav() {
    return <header className={classes.header}>
        <div className = {classes.logo}>
            MHRise Info
        </div>
        <nav>
            <ul>
                <li>
                    <Link to='/'>Home</Link>
                </li>
                <li>
                    <Link to='/armour-sets'>View Armour Sets</Link>
                </li>
            </ul>
        </nav>
    </header>
}

export default MainNav