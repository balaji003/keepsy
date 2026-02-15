import React, { useState } from 'react';
import { View, Text, TextInput, TextInputProps } from 'react-native';

interface InputProps extends TextInputProps {
    label?: string;
    error?: string;
}

export default function Input({ label, error, className, ...props }: InputProps) {
    const [isFocused, setIsFocused] = useState(false);

    return (
        <View className={`space-y-2 ${className}`}>
            {label && <Text className="text-gray-400 font-medium ml-1">{label}</Text>}
            <TextInput
                className={`w-full bg-surface border rounded-xl px-4 py-3.5 text-white placeholder:text-gray-600 transition-all ${isFocused ? 'border-primary shadow-sm shadow-primary/20' : 'border-gray-800'
                    } ${error ? 'border-red-500' : ''}`}
                placeholderTextColor="#6b7280"
                onFocus={() => setIsFocused(true)}
                onBlur={() => setIsFocused(false)}
                {...props}
            />
            {error && <Text className="text-red-500 text-sm ml-1">{error}</Text>}
        </View>
    );
}
