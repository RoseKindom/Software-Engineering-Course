# 作业3：深入理解Callback函数

学号后三位：508

## 要求

参考《代码中的软件工程》第六章可复用软件设计及lab5.2的源代码，完成实验并写一篇实验报告，总结Callback函数的工作机制以及通过参数进行解耦合的方法，深入理解接口设计中的抽象分层。

## Callback 函数

根据维基百科的定义，在计算机程序设计中，回调函数，或简称回调（Callback 即call then back 被主函数调用运算后会返回主函数），是指通过参数将函数传递到其它代码的，某一块可执行代码的引用。这一设计允许了底层代码调用在高层定义的子程序。

## lab5.2 中使用 Callback 的函数模块

在`lab5.2/linktable.c`中，`SearchLinkTableNode` 函数使用了 Callback

```C++
tLinkTableNode * SearchLinkTableNode(tLinkTable *pLinkTable, 
                        int Condition(tLinkTableNode * pNode, void * args),
                        void * args)
{
    if(pLinkTable == NULL || Condition == NULL)
    {
        return NULL;
    }
    tLinkTableNode * pNode = pLinkTable->pHead;
    while(pNode != NULL)
    {    
        if(Condition(pNode, args) == SUCCESS)
        {
            return pNode;				    
        }
        pNode = pNode->pNext;
    }
    return NULL;
}
```

在实现中， `lab5.2/menu.c`中的`FindCmd`函数调用了在同文件中定义的`SearchCOndition`函数，此时`SearchLinkTableNode` 函数中的`Condition`参数即为`SearchCOndition`函数

```C++
int SearchCondition(tLinkTableNode * pLinkTableNode, void * args)
{
    char * cmd = (char*) args;
    tDataNode * pNode = (tDataNode *)pLinkTableNode;
    if(strcmp(pNode->cmd, cmd) == 0)
    {
        return  SUCCESS;  
    }
    return FAILURE;	       
}

/* find a cmd in the linklist and return the datanode pointer */
tDataNode* FindCmd(tLinkTable * head, char * cmd)
{
    return  (tDataNode*)SearchLinkTableNode(head, SearchCondition, (void*)cmd);
}
```

## 耦合

### 耦合定义

耦合性（英语：Coupling）或称耦合力或耦合度，是一种软体度量，是指一程式中，模组及模组之间资讯或参数依赖的程度。

### 耦合分类

耦合性可以是低耦合性（或称为松散耦合），也可以是高耦合性（或称为紧密耦合）。以下列出一些耦合性的分类，从高到低依序排列：

__内容耦合（content coupling，耦合度最高）__
也称为病态耦合（pathological coupling）当一个模块直接使用另一个模块的内部数据，或通过非正常入口而转入另一个模块内部。
__共用耦合/公共耦合（common coupling）__
也称为全局耦合（global coupling.）指通过一个公共数据环境相互作用的那些模块间的耦合。公共耦合的复杂程度随耦合模块的个数增加而增加。
__外部耦合（external coupling）__
发生在二个模组共用一个外加的资料格式、通讯协定或是设备界面，基本上和模组和外部工具及设备的沟通有关。
__控制耦合（control coupling）__
指一个模块调用另一个模块时，传递的是控制变量（如开关、标志等），被调模块通过该控制变量的值有选择地执行块内某一功能;
__特征耦合/标记耦合（stamp coupling）__
也称为资料结构耦合，是指几个模组共享一个复杂的资料结构，如高级语言中的数组名、记录名、文件名等这些名字即标记，其实传递的是这个数据结构的地址;
__资料耦合/数据耦合（data coupling）__
是指模组借由传入值共享资料，每一个资料都是最基本的资料，而且只分享这些资料（例如传递一个整数给计算平方根的函数）。
__讯息耦合（message coupling，是无耦合之外，耦合度最低的耦合）__
可以借由以下二个方式达成：状态的去中心化（例如在物件中），组件间利用传入值或讯息传递来通讯。
__无耦合：模组完全不和其他模组交换资讯。__

## 解耦合

### 定义

在软件工程中，降低耦合度即可以理解为解耦，也就是将强耦合变为弱耦合的过程。

### 为什么要解耦

在软件工程中，对象之间的耦合度就是对象之间的依赖性。对象之间的耦合越高，维护成本越高，因此对象的设计应使类和构件之间的耦合最小，把关联依赖降到最低，而不会产生牵一发而动全身的现象

## lab5.2 中的解耦

`Condition`参数和`SearchCOndition`函数中有一个`args`参数，这个参数的使用属于标记耦合，以下是对它的分析：
`SearchCOndition`函数的功能是判断链表节点中的命令是否与传入的命令相同，满足条件返回 SUCCESS，反之返回 FAILURE。`args`参数在此处是用于传递输入的命令，但同时`args`参数的使用使得查询接口更加通用，配合`Condition`参数将公共耦合降低为标记耦合，从而实现解耦。

