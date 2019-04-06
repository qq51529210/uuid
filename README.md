# uuid

<ul>

<li>
<strong>使用时间戳的版本，uuid.V1()</strong>
<p>默认使用机器MAC地址，可以用SetNode()修改。</p>
<p>MAC地址和和时间戳都决定生成的uuid是否可能出现重复。</p>
</li>

<li>
<strong>使用posix id的版本，uuid.V2()，uuid.V2Uid()，uuid.V2Gid()</strong>
<p>默认使用机器MAC地址系统gid和uid，可以用SetGid()和SetUid()修改</p>
<p>时间戳的前4位置换为id，其他与v1相同。</p>
</li>

<li>
<strong>使用命名空间md5的版本，uuid.V3()</strong>
<p>使用的是md5哈希算法，namespace+name的做哈希。</p>
</li>


<li>
<strong>使用random版本，uuid.V4()</strong>
<p>使用随机数生成的，会有重复的。</p>
</li>

<li>
<strong>使用命名空间sha1的版本，uuid.V5()</strong>
<p>使用的是sha1哈希算法，其他与v3相同。</p>
</li>

<li>
<strong>在分布式下，最后使用V1/V2</strong>
<p>通过设置node的值，不使用MAC，每个uuid服务的就不会一样（保证时间正常）。</p>
</li>

</ul>
