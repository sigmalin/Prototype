using System.Collections;
using System.Collections.Generic;
using System;
using UnityEngine;

namespace Factory
{
    public interface IOrder<T> where T : Enum
    {
        T Type { get; }
    }
}
