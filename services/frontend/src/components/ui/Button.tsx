import { ComponentProps, FC } from "react";
import './button.css'

interface Props extends ComponentProps<"button"> {
    color?: 'primary' | 'secondary' | 'danger';
    size?: 'sm' | 'md' | 'lg';
}

export const Button: FC<Props> = ({color='primary', size='md', ...rest}) => {
    const classes = `main-btn ${color} ${size}`.trimEnd()
    return (
        <>
            <button {...rest} className={classes}/>
        </>
    )
}