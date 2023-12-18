export interface DemoGenTableColumns {    
    id:number;  // ID    
    demoName:string;  // 姓名    
    demoAge:number;  // 年龄    
    classes:string;  // 班级    
    classesTwo:string;  // 班级2    
    demoBorn:string;  // 出生年月    
    demoGender:number;  // 性别    
    createdAt:string;  // 创建日期    
    createdBy:string;  // 创建人    
    demoStatus:number;  // 状态    
    demoCate:string;  // 分类    
    demoThumb:string;  // 头像    
    linkedDemoGenDemoGenClass:LinkedDemoGenDemoGenClass;    
}


export interface DemoGenInfoData {    
    id:number|undefined;        // ID    
    demoName:string|undefined; // 姓名    
    demoAge:number|undefined; // 年龄    
    classes:string|undefined; // 班级    
    linkedClasses:LinkedDemoGenDemoGenClass; // 班级    
    classesTwo:string|undefined; // 班级2    
    linkedClassesTwo:LinkedDemoGenDemoGenClass; // 班级2    
    demoBorn:string|undefined; // 出生年月    
    demoGender:boolean; // 性别    
    createdAt:string|undefined; // 创建日期    
    updatedAt:string|undefined; // 修改日期    
    deletedAt:string|undefined; // 删除日期    
    createdBy:number|undefined; // 创建人    
    updatedBy:number|undefined; // 修改人    
    demoStatus:boolean; // 状态    
    demoCate:any[]; // 分类    
    demoThumb:string|undefined; // 头像    
    demoPhoto:any[]; // 相册    
    demoInfo:string|undefined; // 个人描述    
    demoFile:any[]; // 相关附件    
    linkedDemoGenDemoGenClass:LinkedDemoGenDemoGenClass;    
}


export interface LinkedDemoGenDemoGenClass {	
    id:number|undefined;    // 分类id	
    className:string|undefined;    // 分类名	
}


export interface DemoGenTableDataState {
    ids:any[];
    tableData: {
        data: Array<DemoGenTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            demoName: string|undefined;            
            demoAge: number|undefined;            
            classes: string|undefined;            
            classesTwo: string|undefined;            
            demoBorn: string|undefined;            
            demoGender: number|undefined;            
            createdAt: string|undefined;            
            demoStatus: number|undefined;            
            demoCate: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface DemoGenEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:DemoGenInfoData;
    rules: object;
}