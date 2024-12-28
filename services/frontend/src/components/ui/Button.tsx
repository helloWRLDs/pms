import { ComponentProps, FC, ReactNode } from "react";
import './button.css'
import { IconType } from "react-icons";


type color = 'primary-1' | 'primary-2' | 'primary-3' | 'primary-4' | 'primary-5' | 'primary-3';

interface Props extends ComponentProps<"button"> {
    color?: color;
    size?: 'sm' | 'md' | 'lg';
    icon?: IconType;
    children?: ReactNode;
}

export const Button: FC<Props> = ({color='primary-1', size='md', icon: Icon, children, ...rest}) => {
    const classes = `main-btn ${color} ${size}`.trimEnd()
    return (
        <button {...rest} className={classes}>
            {Icon && <Icon/>}
            {children}
        </button>
    )
}