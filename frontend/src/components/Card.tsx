import React, { useEffect } from 'react';
import { View, ViewProps } from 'react-native';
import Animated, { useAnimatedStyle, useSharedValue, withTiming, withDelay } from 'react-native-reanimated';

interface CardProps extends ViewProps {
    variant?: 'default' | 'elevated';
    delay?: number;
}

export default function Card({ children, className, variant = 'default', delay = 0, ...props }: CardProps) {
    const opacity = useSharedValue(0);
    const translateY = useSharedValue(20);

    useEffect(() => {
        opacity.value = withDelay(delay, withTiming(1, { duration: 500 }));
        translateY.value = withDelay(delay, withTiming(0, { duration: 500 }));
    }, []);

    const animatedStyle = useAnimatedStyle(() => {
        return {
            opacity: opacity.value,
            transform: [{ translateY: translateY.value }],
        };
    });

    const variants = {
        default: "bg-surface border border-gray-800",
        elevated: "bg-surface border border-gray-700 shadow-xl shadow-black/50",
    };

    return (
        <Animated.View
            className={`rounded-2xl p-4 ${variants[variant]} ${className}`}
            style={animatedStyle}
            {...props}
        >
            {children}
        </Animated.View>
    );
}
