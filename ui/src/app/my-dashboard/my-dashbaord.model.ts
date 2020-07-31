  export class SystemInformation {
    public operatingSystem: string;
    public ip: string;
    public memory: Memory;
}

  export class Memory {
    public free: number;
    public total: number;
    public usedPercent: number;
}
