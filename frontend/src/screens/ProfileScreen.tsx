import React from 'react';
import { View, Text, ScrollView, TouchableOpacity, SafeAreaView, StatusBar as RNStatusBar, Switch, Image } from 'react-native';
import { StatusBar } from 'expo-status-bar';
import Card from '../components/Card';
import Button from '../components/Button';

export default function ProfileScreen({ navigation }: any) {
    const [notificationsEnabled, setNotificationsEnabled] = React.useState(true);
    const [darkModeEnabled, setDarkModeEnabled] = React.useState(true);

    const menuItems = [
        { icon: '👤', label: 'Account Settings', action: () => console.log('Account') },
        { icon: '🔒', label: 'Privacy & Security', action: () => console.log('Privacy') },
        { icon: '❓', label: 'Help & Support', action: () => console.log('Help') },
    ];

    return (
        <SafeAreaView className="flex-1 bg-background">
            <StatusBar style="light" />
            <View className="flex-1 px-4 mb-4" style={{ paddingTop: RNStatusBar.currentHeight }}>
                <View className="flex-row justify-between items-center mb-6 pt-4">
                    <TouchableOpacity onPress={() => navigation.goBack()} className="p-2">
                        <Text className="text-white text-lg">Back</Text>
                    </TouchableOpacity>
                    <Text className="text-white text-lg font-bold">Profile</Text>
                    <View className="w-10" />
                </View>

                <ScrollView showsVerticalScrollIndicator={false}>
                    {/* User Profile Header */}
                    <View className="items-center mb-8">
                        <View className="w-24 h-24 bg-primary/20 rounded-full items-center justify-center mb-4 border-2 border-primary">
                            <Text className="text-primary text-4xl font-bold">B</Text>
                        </View>
                        <Text className="text-white text-2xl font-bold">Balaji</Text>
                        <Text className="text-subtext">balaji@example.com</Text>
                    </View>

                    {/* Stats Summary */}
                    <View className="flex-row justify-center space-x-4 mb-8">
                        <Card className="items-center w-28 p-3 bg-surface/50">
                            <Text className="text-white font-bold text-xl">12</Text>
                            <Text className="text-subtext text-xs">Products</Text>
                        </Card>
                        <Card className="items-center w-28 p-3 bg-surface/50">
                            <Text className="text-white font-bold text-xl">5</Text>
                            <Text className="text-subtext text-xs">Bills</Text>
                        </Card>
                    </View>

                    {/* Menu Items */}
                    <Card className="mb-6">
                        {menuItems.map((item, index) => (
                            <View key={item.label}>
                                <TouchableOpacity
                                    className="flex-row items-center py-4"
                                    onPress={item.action}
                                >
                                    <Text className="text-lg mr-4">{item.icon}</Text>
                                    <Text className="text-white text-base flex-1">{item.label}</Text>
                                    <Text className="text-gray-500">›</Text>
                                </TouchableOpacity>
                                {index < menuItems.length - 1 && <View className="h-[1px] bg-gray-800" />}
                            </View>
                        ))}
                    </Card>

                    {/* Preferences */}
                    <Card className="mb-8">
                        <View className="flex-row items-center justify-between py-3">
                            <View classname="flex-row items-center">
                                <Text className="text-lg mr-4">🔔</Text>
                                <Text className="text-white text-base">Notifications</Text>
                            </View>
                            <Switch
                                value={notificationsEnabled}
                                onValueChange={setNotificationsEnabled}
                                trackColor={{ false: '#3e3e3e', true: '#38bdf8' }}
                                thumbColor={notificationsEnabled ? '#ffffff' : '#f4f3f4'}
                            />
                        </View>
                        <View className="h-[1px] bg-gray-800" />
                        <View className="flex-row items-center justify-between py-3">
                            <View classname="flex-row items-center">
                                <Text className="text-lg mr-4">🌙</Text>
                                <Text className="text-white text-base">Dark Mode</Text>
                            </View>
                            <Switch
                                value={darkModeEnabled}
                                onValueChange={setDarkModeEnabled}
                                trackColor={{ false: '#3e3e3e', true: '#38bdf8' }}
                                thumbColor={darkModeEnabled ? '#ffffff' : '#f4f3f4'}
                            />
                        </View>
                    </Card>

                    <Button
                        title="Sign Out"
                        variant="secondary"
                        className="mb-8 border-red-500/50"
                        onPress={() => navigation.replace('Login')}
                    />
                    <Text className="text-red-400 font-bold text-lg text-center absolute bottom-12 left-0 right-0 pointer-events-none">Sign Out</Text>

                </ScrollView>
            </View>
        </SafeAreaView>
    );
}
