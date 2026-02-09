package com.example.starter;




import com.google.common.primitives.Ints;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class OtherTests {
    @Test
    public void testCompare() throws Exception {
        assertEquals(0, Ints.compare(1, 1));
    }

}
