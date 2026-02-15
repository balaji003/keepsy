import React from 'react';
import { Text, View, Pressable, PressableProps } from 'react-native';
import Animated, { useAnimatedStyle, useSharedValue, withSpring } from 'react-native-reanimated';

interface ButtonProps extends PressableProps {
    title: string;
    variant?: 'primary' | 'secondary' | 'outline';
    icon?: React.ReactNode;
    className?: string; // Add className prop explicitly
}

const AnimatedPressable = Animated.createAnimatedComponent(Pressable);

export default function Button({ title, variant = 'primary', icon, className, ...props }: ButtonProps) {
    const scale = useSharedValue(1);

    const animatedStyle = useAnimatedStyle(() => {
        return {
            transform: [{ scale: scale.value }],
        };
    });

    const handlePressIn = () => {
        scale.value = withSpring(0.95);
    };

    const handlePressOut = () => {
        scale.value = withSpring(1);
    };

    const baseClasses = "flex-row items-center justify-center rounded-xl py-4 px-6 transition-all";

    const variants = {
        primary: "bg-primary shadow-lg shadow-primary/30",
        secondary: "bg-surface border border-gray-800",
        outline: "bg-transparent border border-primary",
    };

    const textVariants = {
        primary: "text-white font-bold text-lg",
        secondary: "text-white font-bold text-lg",
        outline: "text-primary font-bold text-lg",
    };

    return (
        <AnimatedPressable
            className={`${baseClasses} ${variants[variant]} ${className}`}
            onPressIn={handlePressIn}
            onPressOut={handlePressOut}
            style={animatedStyle}
            {...props}
        >
            {icon && <View className="mr-2">{icon}</View>}
            <Text className={`${textVariants[variant]}`}>{title}</Text>
        </AnimatedPressable>
    );
}
