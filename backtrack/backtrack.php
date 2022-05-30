<? php

// findTargetSumWays function    leetcode 494 目标和
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    function findTargetSumWays($nums, $target) {
        $result = 0;
        $map = [];

        $fn = function($path, $sum) use(&$nums, &$fn, &$result,$target,&$map){
            if (count($path) == count($nums)) {
                if ($sum == $target) {
                    $result ++;
                    return true;
                }

                return false;
            }

            if (isset($map[$sum."_".count($path)])){
                return false;
            }


            $path[] = -1;
            $b = $fn($path, $sum + $nums[count($path) - 1] * -1);
            array_pop($path);

            $path[] = 1;
            $a = $fn($path, $sum + $nums[count($path) - 1] );
            array_pop($path);

            $ret = $a || $b;
            if ($ret == false) {
                $map[$sum."_".count($path)] = false;
            }

            return $ret;
        };


        $fn([],0);

        return $result;
    }
}
