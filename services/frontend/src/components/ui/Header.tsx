import './ui.css'
import { ComponentProps, FC } from "react";

interface HeaderProps extends ComponentProps<'div'>  {
    logoURL: string;
}

export const Header: FC<HeaderProps> = (props: HeaderProps) => {
    return (
        <header className='header'>
            <a href="/"><img src={props.logoURL} alt="logo" width={100} /></a>
            <ul className='header-list'>
                <li>Home</li>
                <li>About</li>
                <li>Contacts</li>
            </ul>
        </header>
    )
}
